<template>
  <div class="paper">
    <div id="paper_title">
      <li>"{{ obj.title }}"</li>
    </div>
    <div id="paper_authors">
      <li>"{{ obj.authors }}</li>
    </div>
    <div id="paper_equations">
      <macro :macroString="obj.macros"></macro>
      <equation :obj="eq" :key="eq.arxiv_id" v-for="eq in obj.equations" ></equation>
    </div>
  </div>
</template>

<script>
import Equation from './Equation.vue'
import Macro from './Macro.vue'

export default {
  name: 'Paper',
  components: {
    "equation": Equation,
    "macro": Macro,
  },
  props: ['obj'],
  data () {
    return {
      exampleMacro: {
        expression: "\\newcommand{\\bfrac}[2]{\\genfrac{[}{]}{0pt}{}{#1}{#2}}"
      },
    }
  },
}
</script>

<style scoped>
#paper_equations {
  margin-top: 30px;
}
.paper {
  width: 800px;
  margin: 0 auto;
}

li {
  margin: 0 10px;
}
</style>
