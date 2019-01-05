import Vue from 'vue'
import Router from 'vue-router'
import TopPage from '@/components/TopPage'
import FindPaper from '@/components/FindPaper'
import ShowPaper from '@/components/ShowPaper'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'TopPage',
      component: TopPage
    },
  ]
})
