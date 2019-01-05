<template>
  <div class="equation">
    <div class='col'>
      <vue-mathjax class="expression" :formula="mathExp"></vue-mathjax>
    </div>
    <div class='col' id="copy_button">
      <button type="button"
              class='copy_equation'
              v-clipboard:copy="obj.body"
              v-clipboard:success="onCopy"
              v-clipboard:error="onError">Copy
      </button>
    </div>
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
    console.info(this.mathExp)
  },
  methods: {
    onCopy: function (e) {
      console.info('You just copied: ' + e.text)
    },
    onError: function (e) {
      console.info('Failed to copy texts')
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
  display: flex;
  flex-direction: row;
  justify-content: center;
  margin-left: -20px;

  .col {
    margin-left: 20px;
  }
  
  .expression {
    flex: 3; 
  }
  
  #copy_button {
    flex: 1; 
    button {
      height: 20px;
      width: 50px;
    }
  }
}
</style>
