# GM-SDK

#### 介绍
gmsdk 是GMSSH外置应用前端开发所需的sdk，具体文档请查阅 [https://doc-dev.gmssh.com/sdk%E6%8E%A5%E5%85%A5.html](https://doc-dev.gmssh.com/sdk%E6%8E%A5%E5%85%A5.html)}


#### 安装依赖
安装gm-app-sdk 示例
```javascript
npm i gm-app-sdk
//或
yarn add gm-app-sdk
```



#### 接入SDK
在您的入口文件中引入sdk**注意：未避免出现未知错误，请优先在入口文件中的首行引入gm-app-sdk** 示例
```javascript
//main.js的首行
import 'gm-app-sdk';
import { createApp } from 'vue';
```
#### 添加类型
如果您的项目考虑增加ts，可以引入我们的类型文件，并添加到类型扩展中，示例：
```typescript
//global.d.ts 扩展全局 Window 接口
import{ GMProps }from 'gm-app-sdk';

declare global {
    interface Window {
        $gm: GMProps;
    }
}
export {};

```


#### 使用sdk示例
接下来就可以在您的项目中去使用了
```javascript
//使用gm-sdk
window.$gm.openCodeEditor('/www/aaa.txt');
```
