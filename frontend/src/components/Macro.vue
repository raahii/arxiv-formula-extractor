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
  props: ['macros'],
  data()  {
    return {
      "defaultMacros": [
        "\\newcommand{\\bm}[1]{\\boldsymbol #1}",
        "\\newcommand{textnormal}[1]{\\textrm{#1}}",
      ],
    }
  },
  computed: {
    macroExp: function () {
      let exp
      exp  = String.raw`\\(` + "\n"
      exp += this.defaultMacros.join("\n") + "\n"
      exp += this.macros.map(m => m.expression).join("\n") + "\n"
      exp += String.raw`\\)`
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
