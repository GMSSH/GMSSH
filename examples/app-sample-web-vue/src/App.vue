<template>
  <div>
  <button  @click="handleClick1">打开文件系统</button>
  <button  @click="handleClick2">信息提示框</button>
  <button  @click="handleClick3">对话框</button>
  <button  @click="handleClick4">关闭当前应用</button>
  <button  @click="handleClick5">选取文件夹</button>
  <button  @click="handleClick6">选取文件</button>
  <button  @click="handleClick7">获取应用初始尺寸</button>
  <button  @click="handleClick8">打开代码编辑器</button>
  <button  @click="handleClick9">打开终端</button>
  <button  @click="handleClick10">打开日志</button>
  <button  @click="handleClick11">上传文件</button>
</div>
</template>
<script setup>
$gm.childRectListener((res) => {
  const { width, height } = res
  console.log('实时获取应用窗口尺寸，宽：' + width + '，高：' + height);
});
const handleClick1=()=>{
  $gm.openFolder('/');
}
const handleClick2=() => {
  $gm.message.success('打开文件系统成功');
}
const handleClick3=() => {
  $gm.dialog.warning({
    title: '提示',
    content: '这是提示的内容',
    positiveText: '确定',
    negativeText: '取消',
    maskClosable: false,
    onPositiveClick: () => {
      $gm.message.success('点击确定');
      console.log('点击确定');
    },
  });
}

const handleClick4=() => {
  $gm.closeApp();
}

const handleClick5=() => {
  $gm.chooseFolder((res) => {
    $gm.message.success('选择文件夹成功' + res);
    console.log(res);
  }, '/');
}

const handleClick6=() => {
  $gm.chooseFile((res) => {
    $gm.message.success('选择文件成功' + res);
    console.log(res);
  }, '/');
}

const handleClick7=() => {
  const { width, height } = $gm.getRectSize()
  console.log('获取应用初始尺寸，宽：' + width + '，高：' + height);
  $gm.message.success('获取应用初始尺寸，宽：' + width + '，高：' + height);
}

const handleClick8=() => {
  $gm.openCodeEditor('/yarn.lock');
}

const handleClick9=() => {
  $gm.openShell({ arr: ['cd /www\n'] });
}

const handleClick10=() => {
  $gm.openLog('日志名称', {
    fun: () => {
      console.log('日志内容')
    },
    flushed: true,
    control: [
      {
        label: '111确认',
        type: 'negative',
        fun: () => {
          console.log('确认');
        },
      },
    ],
  });
}

const handleClick11=async () => {
  const getWinFiles = async () => {
    // 创建文件选择元素
    const input = document.createElement('input');
    input.setAttribute('type', 'file');
    input.setAttribute('style', 'display:none');
    input.setAttribute('multiple', 'true'); // 多选
    input.setAttribute('accept', ''); // 文件上传格式限制
    // 将文件选择元素添加到文档中
    document.body.appendChild(input);
    // 触发文件选择元素的点击事件
    input.click();
    // 等待文件选择元素的change事件
    return await new Promise((resolve) => {
      input.addEventListener('change', (e) => {
        const files = e.target.files || [];
        // 从文档中移除文件选择元素
        document.body.removeChild(input);
        // 抛出附件，并生成附件uid
        resolve(files);
      });
    });
  }
  const files = await getWinFiles();
  $gm.openUpload({ files, path: '/www' });
}
</script>
<style scoped>

</style>
