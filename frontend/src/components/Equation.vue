<template>
  <div class="equation">
    <katex class='katex col' v-bind:name="obj.id">{{ obj.expression }}</katex>
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
import Katex from './Katex.vue'

export default {
  name: 'Paper',
  components: {
    Katex
  },
  props: ['obj'],
  mounted: function () {
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
  
  .katex {
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
