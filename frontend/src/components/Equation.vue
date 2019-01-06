<template>
  <div class="equation">
    <div class='col'>
      <vue-mathjax class="expression" :formula="obj.expression"></vue-mathjax>
    </div>
    <div class='col' id="copy_button">
      <button type="button"
              class='copy_equation'
              v-clipboard:copy="obj.expression"
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
    console.info(this.obj.expression)
  },
  methods: {
    onCopy: function (e) {
      console.info('You just copied: ' + e.text)
    },
    onError: function (e) {
      console.info('Failed to copy texts')
    }
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
