<template>
  <div class="macro">
    <vue-mathjax :formula="macroExp"></vue-mathjax>
  </div>
</template>

<script>
import { VueMathjax } from 'vue-mathjax'

export default {
  name: 'Macro',
  components: {
   'vue-mathjax': VueMathjax
  },
  mounted: function () {
    console.info("macro:", this.macroExp)
  },
  props: ['macroString'],
  computed: {
    macroExp: function () {
      return "\(" + this.macroString + "\)"
    }
  }
}
</script>

<style scoped lang="scss">
.macro {
  display: none;
}
</style>
