<template>
  <p>{{paper}}</p>
</template>

<script>
import axios from 'axios'

export default {
  name: 'ShowPaper',
  data () {
    return {
      paper: {}
    }
  },
  mounted: function () {
    axios
      .get("http://localhost:1323/papers/"+this.$route.params.paper_id)
      .then(response => {
        if (response.status != 200)  {
          console.log(response.error)
          this.errored = true
          return
        }
        this.paper = response.data.paper
      })
      .catch(error => {
        console.log(error)
        this.errored = true
      })
  },
}
</script>

<style scoped>
</style>
