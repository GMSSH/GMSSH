export const theme = {
  common: {
    // 主色调
    primaryColor: '#5772FF',
    primaryColorHover: '#708BFF',
    primaryColorPressed: '#708BFF',
    primaryColorSuppl: '#148be1',

    // 基础背景色（暗黑风）
    bodyColor: '#101010',
    cardColor: '#101010',
    popoverColor: '#101010',
    modalColor: '#101010',
    tableColor: '#101010',
    // // 文本颜色
    textColorBase: 'var(--jm-accent-7)', // 基础文字色
    textColor1: 'var(--jm-accent-7)', // 主文字色
    textColor2: 'var(--jm-accent-7)', // 次要文字
    textColor3: 'var(--jm-accent-7)', // 禁用/弱提示文字
  },
  Button: {
    // type="primary"
    colorPrimary: 'var(--jm-primary-1)', // 背景色
    textColorPrimary: 'var(--jm-accent-7)',
    textColorHoverPrimary: 'var(--jm-accent-7)',
    textColorPressedPrimary: 'var(--jm-accent-7)',
    textColorFocusPrimary: 'var(--jm-accent-7)',
    colorPressedPrimary: 'var(--jm-primary-1)',
    colorDisabledPrimary: 'var(--jm-primary-1)',
    borderPrimary: '1px solid var(--jm-primary-1)',
    borderHoverPrimary: '1px solid var(--jm-primary-1)',
    borderPressedPrimary: '1px solid var(--jm-primary-1)',
    borderFocusPrimary: '1px solid var(--jm-primary-1)',
    borderDisabledPrimary: '1px solid var(--jm-primary-1)',
    fontWeight: '400',
    // border: '1px solid var(--jm-primary-1)',
    // 小
    heightSmall: '28px',
    fontSizeSmall: '12px',
    paddingSmall: '0 8px',
    // 中
    heightMedium: '34px',
    paddingMedium: '5px 18px',
    fontSizeMedium: '12px',
    // large
    heightLarge: '40px',
    paddingLarge: '9px 26px',
    fontSizeLarge: '14px',

    // waring
    colorWarning: 'var(--jm-primary-1)',
    colorHoverWarning: 'var(--jm-primary-2)',
    colorPressedWarning: 'var(--jm-primary-1)',
    colorFocusWarning: 'var(--jm-primary-2)',
    colorDisabledWarning: 'var(--jm-primary-1)',
    textColorWarning: 'var(--jm-accent-7)',
    textColorHoverWarning: 'var(--jm-accent-7)',
    textColorPressedWarning: 'var(--jm-accent-7)',
    textColorFocusWarning: 'var(--jm-accent-7)',
    textColorDisabledWarning: 'var(--jm-accent-7)',
    borderWarning: '1px solid var(--jm-primary-1)',
    borderPressedWarning: '1px solid var(--jm-primary-1)',
    borderHoverWarning: '1px solid var(--jm-primary-1)',
    borderFocusWarning: '1px solid var(--jm-primary-1)',

    // info
    colorInfo: 'var(--jm-primary-1)',
    colorHoverInfo: 'var(--jm-primary-2)',
    colorPressedInfo: 'var(--jm-primary-1)',
    colorFocusInfo: 'var(--jm-primary-2)',
    colorDisabledInfo: 'var(--jm-primary-1)',
    textColorInfo: 'var(--jm-accent-7)',
    textColorHoverInfo: 'var(--jm-accent-7)',
    textColorPressedInfo: 'var(--jm-accent-7)',
    textColorFocusInfo: 'var(--jm-accent-7)',
    textColorDisabledInfo: 'var(--jm-accent-7)',
    borderInfo: '1px solid var(--jm-primary-1)',
    borderPressedInfo: '1px solid var(--jm-primary-1)',
    borderHoverInfo: '1px solid var(--jm-primary-1)',
    borderFocusInfo: '1px solid var(--jm-primary-1)',

    // Success
    colorSuccess: 'var(--jm-primary-1)',
    colorHoverSuccess: 'var(--jm-primary-2)',
    colorPressedSuccess: 'var(--jm-primary-1)',
    colorFocusSuccess: 'var(--jm-primary-2)',
    colorDisabledSuccess: 'var(--jm-accent-3)',
    textColorSuccess: 'var(--jm-accent-7)',
    textColorHoverSuccess: 'var(--jm-accent-7)',
    textColorPressedSuccess: 'var(--jm-accent-7)',
    textColorFocusSuccess: 'var(--jm-accent-7)',
    textColorDisabledSuccess: 'var(--jm-accent-7)',
    borderSuccess: '1px solid var(--jm-primary-1)',
    borderPressedSuccess: '1px solid var(--jm-primary-1)',
    borderHoverSuccess: '1px solid var(--jm-primary-1)',
    borderFocusSuccess: '1px solid var(--jm-primary-1)',

    // error
    colorError: 'var(--jm-primary-1)',
    colorHoverError: 'var(--jm-primary-2)',
    colorPressedError: 'var(--jm-primary-1)',
    colorFocusError: 'var(--jm-primary-2)',
    colorDisabledError: 'var(--jm-accent-3)',
    textColorError: 'var(--jm-accent-7)',
    textColorHoverError: 'var(--jm-accent-7)',
    textColorPressedError: 'var(--jm-accent-7)',
    textColorFocusError: 'var(--jm-accent-7)',
    textColorDisabledError: 'var(--jm-accent-7)',
    borderError: '1px solid var(--jm-primary-1)',
    borderPressedError: '1px solid var(--jm-primary-1)',
    borderHoverError: '1px solid var(--jm-primary-1)',
    borderFocusError: '1px solid var(--jm-primary-1)',

    // Disabled
    textColorDisabled: 'var(--jm-accent-7)',
    textColorDisabledPrimary: 'var(--jm-accent-7)',
  },
  Select: {
    menuBoxShadow: '0 0 1px var(--jm-accent-5)',
    peers: {
      InternalSelection: {
        fontWeight: '400',
        fontSizeMedium: '13px',
        placeholderColor: 'var(--jm-accent-4)',
        borderRadius: '4px',
        color: 'var(--jm-theme)',
        textColor: 'var(--jm-accent-7)',
        border: '1px solid var(--jm-accent-2)',
        paddingSingle: '10px',
      },
      InternalSelectMenu: {
        fontWeight: '400',
        optionFontSizeMedium: '13px',
        optionHeightMedium: '34px',
        color: 'var(--jm-accent-1)',
        optionTextColorHover: 'var(--jm-accent-7)',
        optionTextColor: 'var(--jm-accent-7)',
        border: '1px solid var(--jm-accent-2)',
        borderRadius: '4px', // 可选：圆角
      },
    },
  },
  Input: {
    heightMedium: '34px', // 输入框高度
    fontSizeMedium: '12px', // 字体大小
    borderRadius: '4px', // 圆角
    border: '1px solid var(--jm-accent-3)', // 边框样式
    color: 'var(--jm-theme)', // 背景色
    textColor: 'var(--jm-accent-7)', // 文本颜色
    placeholderColor: 'var(--jm-accent-4)', // placeholder 颜色
    caretColor: 'var(--jm-primary-1)', // 光标颜色
    colorFocus: 'var(--jm-theme)',
  },
  AutoComplete: {
    menuBoxShadow: '0 0 2px var(--jm-accent-5)',
    peers: {
      InternalSelectMenu: {
        color: 'var(--jm-theme)', // ✅ 下拉框背景色
        border: '1px solid var(--jm-accent-3)', // 边框样式
        optionColorHover: 'var(--jm-primary-3)', // ✅ 悬浮时背景
        optionTextColor: 'var(--jm-primary-1)', // ✅ 文字颜色
        menuBoxShadow: '1px 2px 12px var(--jm-accent-7)',
      },
    },
  },
  Checkbox: {
    checkMarkColor: 'var(--jm-accent-7)',
  },
  Radio: {
    dotColor: 'var(--jm-primary-1)',
    textColor: 'var(--jm-accent-7)',
    buttonBorderColor: 'var(--jm-accent-3)',
    buttonBorderColorActive: 'var(--jm-primary-1)',
    buttonColorActive: 'var(--jm-primary-1)',
    buttonTextColor: 'var(--jm-accent-7)',
    buttonTextColorActive: 'var(--jm-accent-7)',
  },
  Tooltip: {
    color: 'rgba(57,57,57,0.95)',
    padding: '4px 8px',
    borderRadius: '4px',
    boxShadow: '0 2px 4px 0 rgba(0, 0, 0, 0.25)',
  },
  Popover: {
    fontSize: '13px',
  },
  // Dialog
  Dialog: {
    color: 'var(--jm-theme)',
    iconColorError: 'var(--jm-error-color)',
    iconColorWarning: 'var(--jm-warning-color)',
    iconColorSuccess: 'var(--jm-success-color)',
    iconColorInfo: 'var(--jm-primary-1)',
  },
  Message: {
    iconColorSuccess: 'var(--jm-success-color)',
    iconColorWarning: 'var(--jm-warning-color)',
    iconColorError: 'var(--jm-error-color)',
    iconColorInfo: 'var(--jm-primary-1)',
    borderRadius: '8px',
    textColor: 'var(--jm-primary-1)',
  },
  DatePicker: {
    panelColor: 'transparent',
    calendarTitleTextColor: 'var(--jm-accent-7)',
    arrowColor: 'var(--jm-accent-7)',
    itemTextColor: 'var(--jm-accent-7)',
    itemColorHover: 'var(--jm-primary-3)',
    itemColorActive: 'var(--jm-primary-1)',
    itemTextColorActive: 'var(--jm-accent-7)',
  },
  DataTable: {
    tdColor: 'var(--jm-accent-1)',
    thColor: 'var(--jm-accent-1)',
    borderColor: 'var(--jm-theme)',
    fontSizeSmall: '12px',
    fontSizeMedium: '12px',
    thPaddingMedium: '10px',
    tdPaddingMedium: '10px',
    // borderRadius: '4px',
  },
  Card: {
    color: 'var(--jm-theme)',
    borderColor: 'var(--jm-accent-2)',
    textColor: 'var(--jm-accent-7)',
    titleTextColor: 'var(--jm-accent-7)',
  },
  Modal: {
    color: 'var(--jm-theme)',
    textColor: 'var(--jm-accent-7)',
    boxShadow: '0 0 20px rgba(0,0,0,0.5)',
  },
  Drawer: {
    color: 'var(--jm-theme)',
    textColor: 'var(--jm-accent-7)',
    boxShadow: '-2px 0 8px rgba(0,0,0,0.2)',
  },
  Switch: {
    railColor: 'var(--jm-accent-3)',
    railColorActive: 'var(--jm-primary-1)',
    loadingColor: 'var(--jm-primary-1)',
  },
  Dropdown: {
    color: 'var(--jm-theme)',
    optionTextColor: 'var(--jm-accent-7)',
    dividerColor: 'var(--jm-accent-2)',
  },
  Popconfirm: {
    color: 'var(--jm-theme)',
    textColor: 'var(--jm-accent-7)',
  },
  Skeleton: {
    color: 'var(--jm-accent-2)',
  },
  Tag: {
    color: 'var(--jm-accent-1)',
    textColor: 'var(--jm-accent-7)',
    borderColor: 'var(--jm-accent-3)',
  },
};
