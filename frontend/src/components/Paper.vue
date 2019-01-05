<template>
  <div class="paper">
    <div id="paper_title">
      <li>"{{ obj.title }}"</li>
    </div>
    <div id="paper_authors">
      <li>"{{ obj.authors }}</li>
    </div>
    <div id="paper_equations" v-for="eq in obj.equations">
      <equation v-bind:obj="eq"></equation>
    </div>
  </div>
</template>

<script>
import Equation from './Equation.vue'

export default {
  name: 'Paper',
  components: {
    Equation
  },
  props: ['obj'],
  data () {
    return {
    }
  },
  mounted: function () {
    console.log(this.obj.equations)
  },
  methods: {
  }
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
