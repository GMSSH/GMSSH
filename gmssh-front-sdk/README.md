# GM-SDK

#### Introduction
gmsdk is the SDK required for front-end development of GMSSH external applications. For detailed documentation, please refer to [https://doc-dev.gmssh.com/sdk%E6%8E%A5%E5%85%A5.html](https://doc-dev.gmssh.com/sdk%E6%8E%A5%E5%85%A5.html)


#### Installation
Install gm-app-sdk example:
```javascript
npm i gm-app-sdk
// or
yarn add gm-app-sdk
```



#### Integrate SDK
Import the SDK in your entry file. **Note: To avoid unknown errors, please import gm-app-sdk at the first line of your entry file.** Example:
```javascript
// First line of main.js
import 'gm-app-sdk';
import { createApp } from 'vue';
```
#### Add Types
If your project uses TypeScript, you can import our type definitions and add them to the type extensions. Example:
```typescript
// global.d.ts - Extend the global Window interface
import { GMProps } from 'gm-app-sdk';

declare global {
    interface Window {
        $gm: GMProps;
    }
}
export {};

```


#### SDK Usage Example
Now you can use it in your project:
```javascript
// Using gm-sdk
window.$gm.openCodeEditor('/www/aaa.txt');
```
