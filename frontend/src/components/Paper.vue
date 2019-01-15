<template>
  <div class="paper">
    <div class="paper_title">
      <a :href="obj.url" target="_blank">{{ obj.title }}</a>
    </div>
    <div class="paper_authors">
      {{ this.authorsStr }}
    </div>
    <div class="paper_equations">
      <p class="n_hits success"> {{ this.obj.equations.length }} equations found.</p>
      <macro :macros="obj.macros"></macro>
      <equation :eq="eq" :macros="equationMacros[i]" :key="eq.id" v-for="(eq, i) in obj.equations" ></equation>
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
    }
  },
  computed: {
    authorsStr: function () {
      return this.obj.authors.map(a => a.name).join(", ")
    },
    equationMacros: function () {
      let macros = []
      for (let eq of this.obj.equations) {
        let _macros = []
        for (let m of this.obj.macros) {
          if (eq.expression.indexOf(m.command) >= 0) {
            _macros.push(m)
          }
        }
        macros.push(_macros)
      }
      return macros
    },
  },
}
</script>

<style scoped lang="scss">
#paper_equations {
  margin-top: 30px;
}
.paper {
  width: 100%;
  margin: 40px auto;

  .paper_title {
    margin: 5px 0px;
    font-size: x-large;
    font-weight: bold;

    a {
      color: #2c3e50;
      text-decoration: none;
    }
  }

  .paper_authors {
    margin: 5px 0px;
    font-size: medium;
    line-height: 150%;
  }
  .paper_equations {
    .n_hits {
      color: blue;
    }
  }
}

li {
  margin: 0 10px;
}
</style>
