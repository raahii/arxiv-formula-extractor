<template>
  <div class="equation card">
    <vue-mathjax 
      class="expression" 
      :formula="mathExp"
      v-clipboard:copy="obj.expression"
      v-clipboard:success="onCopy"
      v-clipboard:error="onError"></vue-mathjax>
  </div>
</template>

<script>
import { VueMathjax } from 'vue-mathjax'

export default {
  name: 'Paper',
  components: {
   'vue-mathjax': VueMathjax
  },
  props: ['obj'],
  mounted: function () {
    // console.info(this.mathExp)
  },
  methods: {
    onCopy: function (e) {
      // console.info('You just copied: ' + e.text)
    },
    onError: function (e) {
      // console.info('Failed to copy texts')
    }
  },
  computed: {
    mathExp: function () {
      let exp;
      exp  = String.raw`\begin{eqnarray}`
      exp += this.obj.expression
      exp += String.raw`\end{eqnarray}`
      return exp
    },
  }
}
</script>

<style scoped lang="scss">
.equation {
  width: 100%;
  margin-bottom: 20px;
  padding: 20px 0;
}

.card {
  border: 1px solid;
  border-color: #cccccc;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
  transition: all 0.3s cubic-bezier(.25,.8,.25,1);
}

.card:hover {
  cursor: pointer;
  box-shadow: 0 5px 5px rgba(0,0,0,0.25), 0 5px 5px rgba(0,0,0,0.22);
}
</style>
