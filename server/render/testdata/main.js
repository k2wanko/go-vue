import Vue from 'vue'
import { createRenderer } from 'vue-server-renderer'

const app = new Vue({
    render: function (h) {
        return h('p', 'hello world')
    }
})

const layout = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>testdata</title>
  </head>
  <body>
    <div id="app"></div>
  </body>
</html>
`

const html = (() => {
    const target = '<div id="app"></div>'
    const i = layout.indexOf(target)
    return {
        head: layout.slice(0, i),
        tail: layout.slice(i + target.length)
    }
})()

const renderer = createRenderer({
    cache: global['__ComponentCache__'],
})

module.exports = function (context) {
    const res = context.res
    const stream = renderer.renderToStream(app)
    stream.once('data', () => {
        res.write(html.head)
    })
    stream.on('data', chunk => {
        res.write(chunk)
    })
    stream.on('end', () => {
        res.end(html.tail)
    })
    stream.on('error', err => {
        res.error(err)
    })
}