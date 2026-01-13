(function () {
    // 设置当前域名 不包含二级域名
    // 判断是否为IP地址，如果不是IP地址才设置document.domain
    const hostname = window.location.hostname;
    const isIP = /^(\d{1,3}\.){3}\d{1,3}$/.test(hostname);
    if (!isIP) {
        try {
            document.domain = hostname.split('.').slice(-2).join('.');
        } catch (error) {
            console.warn('无法设置顶级域名，默认设置为当前域名', error);
        }
    }
    // 私有状态存储
    const state = {
        resizePadding: 6,
        width: 0,
        height: 0,
        listeners: new Set(),
        gmcCallback: null,
        serveActiveCallback: null,
        extAppStatusCallback: null,
    };
    //  获取url参数
    const getURLParams = (url) => {
        // 如果提供了url参数，使用它；否则使用当前页面的URL
        const targetUrl = url || window.location.href;

        // 尝试从查询字符串获取参数
        const urlObj = new URL(targetUrl);
        const searchParams = new URLSearchParams(urlObj.search);
        const params = {};

        // 获取查询字符串参数
        for (const [key, value] of searchParams) {
            params[key] = value;
        }

        // 如果查询字符串中没有参数，尝试从hash中获取
        if (Object.keys(params).length === 0 && urlObj.hash) {
            const hashParts = urlObj.hash.split('?');
            if (hashParts.length > 1) {
                const hashParams = new URLSearchParams(hashParts[1]);
                for (const [key, value] of hashParams) {
                    params[key] = value;
                }
            }
        }

        return params;
    }
    // 初始化 $gm 对象
    window.$gm = { ...window.parent.$gm, ...getURLParams(window.location.href) } || {};
    // 消息类型处理器映射
    const messageHandlers = {
        // init: (payload) => {
        //   window.$gm={...window.$gm,...payload}
        // },
        resize: (payload) => {
            state.width = payload.width;
            state.height = payload.height;
            state.listeners.forEach((fn) =>
                fn({ width: payload.width, height: payload.height }),
            );
        },
        gmcListener: (payload) => {
            if (state.gmcCallback) {
                state.gmcCallback(payload);
            }
        },
        serveListener: (payload) => {
            if (state.serveActiveCallback) {
                state.serveActiveCallback(payload);
            }
        },
        extAppStatus: (payload) => {
            if (state.extAppStatusCallback) {
                state.extAppStatusCallback(payload);
            }
        },
    };

    // 统一消息处理函数
    function handleMessageEvent(e) {
        try {
            const data = e.data || {};
            if (!data.type || !messageHandlers[data.type]) {
                return;
            }

            messageHandlers[data.type](data.data);
        } catch (error) {
            console.error('Message handling error:', error);
        }
    }

    // 注册全局消息监听器（仅此一个）
    window.addEventListener('message', handleMessageEvent);

    // ====== API 接口实现 ======
    // 监听 resize 消息
    window.$gm.childRectListener = (callback) => {
        if (callback) {
            state.listeners.add(callback);
        }
        return () => callback && state.listeners.delete(callback);
    };

    // 监听 GMC 消息
    window.$gm.mainGMCListener = (callback) => {
        state.gmcCallback = callback;
    };

    // 监听 GMC 消息
    window.$gm.serveActiveListener = (callback) => {
        state.serveActiveCallback = callback;
    };

    // 监听 extAppStatus 消息
    window.$gm.extAppStatusListener = (callback) => {
        state.extAppStatusCallback = callback;
    };
    // ====== 保持不变的鼠标监听逻辑 ======
    document.onmousemove = (e) => {
        const inResizeZone =
            e.clientX < state.resizePadding ||
            e.clientY < state.resizePadding ||
            e.clientX > window.innerWidth - state.resizePadding ||
            e.clientY > window.innerHeight - state.resizePadding;

        const currentState = inResizeZone ? 'resizeStart' : 'resizeEnd';
        window.parent.postMessage(
            { type: currentState, data: window.$gm.fileId },
            '*',
        );
    };

    document.addEventListener('mousedown', () => {
        window.parent.postMessage(
            { type: 'iframeMouseDown', data: window.$gm.fileId },
            '*',
        );
    })
    // ====== 工具方法 ======
    // 获取窗口尺寸
    window.$gm.getRectSize = () => ({
        width: state.width,
        height: state.height,
    });

    // 关闭应用方法
    window.$gm.closeApp = function () {
        if (window.$gm.fileId) {
            window.parent.postMessage(
                { type: 'closeApp', data: window.$gm.fileId },
                '*',
            );
        }
    };

    // 发送消息至父级
    window.$gm.emitParent = function (msg) {
        window.parent.postMessage(
            msg,
            '*',
        );
    };

    // 应用初始化方法
    window.$gm.init = function () {
        return window.$gm
            .request('/api/center/check_status', {
                method: 'post',
                data: {
                    app_name: window?.$gm?.name,
                    version: window?.$gm?.version,
                    communication_type: window?.$gm?.communicationType,
                },
            })
    };

    // 设置应用窗口样式
    window.$gm.setAppRectStyle = function (style = {}) {
        if (window.$gm.fileId) {
            window.parent.postMessage(
                { type: 'setAppRectStyle', data: window.$gm.fileId, style },
                '*',
            );
        }
    };

    // 设置主题
    const style = document.createElement('style');
    style.textContent = window.$gm.themeCss || '';
    document.head.appendChild(style);
    document.documentElement.setAttribute('data-theme', 'dark');
})();