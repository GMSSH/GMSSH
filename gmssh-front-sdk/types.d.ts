export interface GMProps {
  /**
   * 当前登录连接的token
   */
  token: string;
  /**
   * 主应用传递过来的自定义数据
   */
  data: any;
  /**
   * 当前应用的版本号
   */
  version: string;
  /**
   * 每个应用启动时，会自动注入一个fileId
   */
  fileId: string;

  communicationType: string;
  /**
   * 当前应用的名称
   */
  name: string;
  /**
   * 主应用的当前url
   */
  webURL: string;
  /**
   * 当前连接到服务器host
   */
  host: string;
  /**
   * 当前的语言类型，和主应用同步
   */
  lang: string;

  /**
   * 信息提示，会调用主应用的message API
   */
  message: any;
  /**
   * 弹出对话框，会调用主应用的dialog API
   */
  dialog: any;
  /**
   * 关闭当前窗口
   */
  closeApp: () => void;
  /**
   * 关闭对应fileId的应用，会触发childDestroyedListener事件的回调
   * @param fileId
   */
  closeWindow: (fileId: string) => void;
  /**
   * 选取文件夹，参数1为回调函数，返回选择的路径，参数2为默认打开的路径，默认为'/'
   * @param callback
   * @param path
   */
  chooseFolder: (callback: (v: string) => void, path?: string) => void;
  /**
   * 选取文件，参数1为回调函数，返回选择的路径，参数2为默认打开的路径，默认为'/'
   * @param callback
   * @param path
   */
  chooseFile: (callback: (v: string) => void, path?: string) => void;

  /**
   * 获取当前应用尺寸信息，通常用于初始化时获取
   * @returns {width:number;height:number}
   */
  getRectSize: () => { width: number; height: number };

  /**
   * 监听主应用对单个子应用的通知消息推送
   * @param callback
   */
  mainNotificationListener: (callback: (v: any) => void) => void;
  /**
   * 监听主应用对单个子应用的GMC通知消息推送
   * @param callback
   */
  mainGMCListener: (callback: (v: any) => void) => void;
  /**
   * 监听子应用尺寸变化
   * @param callback
   */
  childRectListener: (
    callback: (v: { width: number; height: number }) => void
  ) => void;
  /**
   * 打开文件系统，path为默认打开的文件夹路径
   * @param path
   */
  openFolder: (path: string) => void;
  /**
   * 打开代码编辑器，path为文件所在路径
   * @param path
   */
  openCodeEditor: (path: string) => void;
  /**
   * 打开终端，arr为可执行命令数组
   * @param data
   */
  openShell: (data: { arr: string[] }) => void;
  /**
   * 上传文件到服务器 files为文件对象，可配合input上传使用，path为上传对应服务器文件夹的路径
   * @param files
   * @param string
   */
  openUpload: ({ files: files, path: string }) => void;
  /**
   * 渲染一个自定义窗口，注：render函数需要特殊支持，如vue框架等
   * @param render
   * @param string
   */
  renderCustomWindow: ({ render: render, title: string }) => void;
  /**
   * 打开一个子应用窗口
   * @param data
   */
  renderChildWindow: (data: {
    name: string;
    url: string;
    mini: boolean;
    data: any;
    width: number;
  }) => void;

  request<T>(s: string, param2: any): Promise<T>;

  naiveUiTheme: { [key: string]: string };
  baseURL: string;

  downloadFile: (url: string) => void;
  isExperiencesServer: boolean;
  setDesktopWallpaper: (url: string) => Promise<void>;
  init: () => Promise<void>;
  modalClose: () => void;
}
