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
      let exp
      exp  = String.raw`\\(` + "\n"
      exp += this.macroString
      exp += "\n" + String.raw`\\)`
      return exp
    }
  }
}
</script>

<style scoped lang="scss">
.macro {
  display: none;
}
</style>
