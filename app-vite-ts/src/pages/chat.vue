<template>
  <view class="container">
    <uni-list :border="true">
      <!-- 头像显示圆点 -->
      <uni-list-chat
        title="标题 uni-app 题 uni-app题 uni-app题 uni-app题 uni-app"
        avatar="https://web-assets.dcloud.net.cn/unidoc/zh/unicloudlogo.png"
        note="您收到一条新的消息您收到一条新的消息您收到一条新的消息您收到一条新的消息您收到一条新的"
        time="02-02 20:20"
        to="detail/chat/index"
      ></uni-list-chat>
      <uni-list-chat
        title="标题 uni-app 题 uni-app题 uni-app题 uni-app题 uni-app"
        avatar="https://web-assets.dcloud.net.cn/unidoc/zh/unicloudlogo.png"
        note="您收到一条新的消息您收到一条新的消息您收到一条新的消息您收到一条新的消息您收到一条新的"
        time="02-02 20:20"
      ></uni-list-chat>
      <uni-list-chat
        title="标题 uni-app 题 uni-app题 uni-app题 uni-app题 uni-app"
        avatar="https://web-assets.dcloud.net.cn/unidoc/zh/unicloudlogo.png"
        note="您收到一条新的消息您收到一条新的消息您收到一条新的消息您收到一条新的消息您收到一条新的"
        time="02-02 20:20"
      ></uni-list-chat>
    </uni-list>
  </view>
</template>
<script setup lang="ts">
function goDetail() {
  uni.navigateTo({ url: 'detail/chat/index' })
}
</script>

<style lang="scss" scoped>
.container {
  .slot-image {
    height: 50px;
    width: 50px;
  }
}
</style>
