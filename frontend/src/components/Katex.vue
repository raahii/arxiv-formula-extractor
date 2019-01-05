<template>
  <div :id="name">
    <slot>...</slot>
  </div>
</template>

<script>
export default {
  props: ['name'],
  mounted() {
    const element = document.getElementById(this.name)
    katex.render(this.$slots.default[0].text, element)
  }
}
</script>

<style scoped lang="scss">
div {
  text-align: center;
}
</style>
