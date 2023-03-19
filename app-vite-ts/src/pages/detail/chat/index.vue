<template>
  <!-- #ifdef H5 -->
  <view class="chat-container max-w-6xl" ref="chatRef">
    <view class="list">
      <view class="item" v-for="(item, index) in messageList" :key="index">
        <view class="nick" v-if="item.role !== 'user'" :class="[item.role]">
          <image class="avatar" src="https://web-assets.dcloud.net.cn/unidoc/zh/unicloudlogo.png" mode="aspectFill" />
          <text>AI助手</text>
        </view>
        <view class="content-wrapper" :class="[item.role, item.role == 'assistant' ? 'mini' : '']">
          <view class="content" :class="[item.role, reqStatus ? 'req-status' : '']">
            <!-- {{ item.content }} -->
            <view v-if="item.role == 'user'"> {{ item.content }}</view>
            <view v-else v-html="mdi.render(item.content)" class="markdown-body"></view>
          </view>
          <image v-if="item.role == 'user'" src="https://web-assets.dcloud.net.cn/unidoc/zh/unicloudlogo.png" mode="aspectFill" class="avatar" />
        </view>
      </view>
    </view>
    <view class="bottom">
      <view class="max-w-5xl flex m-auto">
        <input class="input" type="text" v-model="messageInput" placeholder="输入你想说的话" confirm-type="search" @confirm="sendMessage" />
        <button class="button" @click="sendMessage">发送</button>
      </view>
    </view>
  </view>
  <!-- #endif -->
  <!-- #ifdef MP-WEIXIN -->
  <web-view src="http://192.168.1.101:6001/" />
  <!-- #endif -->
</template>
<script setup lang="ts">
import MarkdownIt from 'markdown-it'
import { marked } from 'marked'
import mdKatex from '@traptitech/markdown-it-katex'
import hljs from 'highlight.js'
import 'highlight.js/styles/Dark.css'

const mdi = new MarkdownIt({
  linkify: true,
  highlight(code: any, language: any) {
    const validLang = !!(language && hljs.getLanguage(language))
    if (validLang) {
      const lang = language ?? ''
      return highlightBlock(hljs.highlight(lang, code, true).value, lang)
    }
    return highlightBlock(hljs.highlightAuto(code).value, '')
  },
})

mdi.use(mdKatex, { blockClass: 'katexmath-block rounded-md p-[10px]', errorColor: ' #cc0000' })

// const render = new marked.Renderer()
// marked.setOptions({
//   renderer: render, // 这是必填项
//   gfm: true, // 启动类似于Github样式的Markdown语法
//   pedantic: false, // 只解析符合Markdwon定义的，不修正Markdown的错误
//   sanitize: false, // 原始输出，忽略HTML标签（关闭后，可直接渲染HTML标签）

//   // 高亮的语法规范
//   highlight: (code, lang) => hljs.highlight(code, { language: lang }).value,
// })

const reqStatus = ref(false)
let stream: any = null
let timer: any = null
const messageAll = ref('')
const messageInput = ref('')
const messageList: any = reactive([])

// const text = computed(() => {
//     return 
// })
onMounted(() => {})
// function textAuto(str: string) {
//   messageAll.value += str
//   clearInterval(timer)
//   timer = setInterval(() => {
//     console.log('messageAll.value', messageList[messageList.length - 1].content)
//     if (messageAll.value.length > messageList[messageList.length - 1].content.length) {
//       messageList[messageList.length - 1].content += messageAll.value.charAt(messageList[messageList.length - 1].content.length)
//     } else {
//       console.log('ttimer = null')
//       clearInterval(timer)
//     }
//   }, 50)
// }
async function sendMessage() {
  const messageParams = messageList.slice(-5)
  messageList.push({ role: 'user', content: messageInput.value }, { role: 'assistant', content: '' })

  const resquestOptions = {
    body: JSON.stringify({ messages: [...messageParams, { role: 'user', content: messageInput.value }] }),
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
  }
  reqStatus.value = true
  messageInput.value = ''
  fetch('http://localhost/app/getGpt', resquestOptions)
    .then(async (response) => {
      const reader = await response.body?.getReader()
      while (true) {
        const obj: any = await reader?.read()
        console.log('obj', obj)

        if (obj.done) {
          console.log('done', obj.done)
          let timer = setTimeout(() => {
            reqStatus.value = false
            clearTimeout(timer)
          }, 500)
          break
        }
        // const t = new TextDecoder('utf-8').decode(new Uint8Array(obj.value))
        // console.log('t', t)

        const t = new TextDecoder('utf-8').decode(new Uint8Array(obj.value)).split('data:')

        const contentArr = t.map((item) => {
          try {
            const obj = JSON.parse(item)
            //显示消息
            messageList[messageList.length - 1].content += obj.content
            scrollBottom()
            // messageAll.value += obj.content
            // textAuto(obj.content)
            console.log('content', { content: obj.content })
          } catch (error) {
            return '' // 处理解析失败的情况
          }
        })
      }
    })
    .catch((error) => {
      console.error('Error:', error)
    })
}
function scrollBottom() {
  nextTick(() => {
    const chatc: any = document.querySelector('.chat-container')
    console.log('chatc.scrollHeight - 800', chatc.scrollHeight - 800)
    chatc.scroll({
      left: 0,
      top: chatc.scrollHeight - 90,
      behavior: 'smooth',
    })
  })
}
function highlightBlock(str: string, lang?: string) {
  return `<pre class="code-block-wrapper"><div class="code-block-header"><span class="code-block-header__lang">${lang}</span><span class="code-block-header__copy"></span></div><code class="hljs code-block-body ${lang}">${str}</code></pre>`
}
</script>

<style lang="scss" scoped>
.chat-container {
  background: #f2f2f2;
  padding: 14px 14px 90px;
  height: calc(100vh - 90px);
  box-sizing: border-box;
  overflow: auto;
  margin: auto;
  .revoked {
    text-align: center;
    font-size: 12px;
    color: #999;
  }
  .bottom {
    position: fixed;
    left: 0;
    bottom: 0;
    padding: 10px 10px 30px;
    right: 0;
    // display: flex;
    // align-items: center;
    // justify-content: space-between;
    background-color: #fff;
  }
  .button {
    width: 70px;
    padding: 0;
    margin: 0;
    line-height: 35px;
    height: 35px;
    background-color: #07c160;
    color: #fff;
    font-size: 14px;
  }
  .button:after {
    padding: 0;
    margin: 0;
    border: none;
  }
  .input {
    border: 1px solid #eee;
    height: 35px;
    padding: 0 10px;
    box-sizing: border-box;
    border-radius: 4px;
    flex: 1;
    color: #333;
    margin-right: 10px;
    font-size: 14px;
  }
  .list .item:last-child .content-wrapper .content.req-status:after {
    -webkit-animation: blink 1s steps(5, start) infinite;
    animation: blink 1s steps(5, start) infinite;
    content: '▋';
    vertical-align: baseline;
  }
  @keyframes blink {
    from {
      opacity: 0;
    }

    50% {
      opacity: 1;
    }

    100% {
      opacity: 0;
    }
  }
  .item {
    color: #333;
    margin-bottom: 12px;
    word-break: break-all;
    display: flex;
    flex-direction: column;

    .avatar {
      width: 30px;
      height: 30px;
      border-radius: 50%;
      margin-right: 6px;
      flex-shrink: 0;
    }
    .nick {
      color: #333;
      font-size: 12px;
      display: flex;
      align-items: center;
      margin-bottom: 6px;
    }
    .nick.user {
      align-self: flex-end;
    }
    .content-wrapper {
      display: flex;
      align-items: flex-start;
      justify-content: flex-end;
      .content {
        padding: 6px 10px;
        background-color: rgb(149, 236, 105);
        border-radius: 4px;
        margin-left: 34px;
        // margin-right: 34px;
        // flex: 1;
      }
      .content.assistant {
        // margin-left: 0;
        margin-right: 34px;
        background-color: #fff;
        color: #333;
      }
      .content.user {
        margin-right: 10px;
      }
    }
    .content-wrapper.assistant {
      justify-content: flex-start !important;
    }
  }
  // pre {
  //   background-color: #f5f5f5;
  //   padding: 10px;
  //   border-radius: 5px;
  // }
  /* #ifdef MP-WEIXIN */
  .mini {
    display: flex;
    justify-content: flex-start !important;
    .content {
      margin-right: 34px;
      background-color: #fff !important;
      color: #333;
    }
  }
  /* #endif */
}
.markdown-body {
	background-color: transparent;
	font-size: 14px;

	p {
		white-space: pre-wrap;
	}

	ol {
		list-style-type: decimal;
	}

	ul {
		list-style-type: disc;
	}

	pre code,
	pre tt {
		line-height: 1.65;
	}

	.highlight pre,
	pre {
		background-color: #fff;
	}

	code.hljs {
		padding: 0;
	}

	.code-block {
		.code-block-wrapper {
			position: relative;
			padding-top: 24px;
		}

		.code-block-header {
			position: absolute;
			top: 5px;
			right: 0;
			width: 100%;
			padding: 0 1rem;
			display: flex;
			justify-content: flex-end;
			align-items: center;
			color: #b3b3b3;

			.code-block__copy{
				cursor: pointer;
				margin-left: 0.5rem;
				user-select: none;
				&:hover {
					color: #65a665;
				}
			}
		}
	}
}

html.dark {

	.highlight pre,
	pre {
		background-color: #282c34;
	}
}
</style>
