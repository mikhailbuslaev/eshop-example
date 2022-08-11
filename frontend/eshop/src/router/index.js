import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import AboutView from '../views/AboutView.vue'
import CardView from '../views/CardView.vue'
const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/about',
    name: 'about',
    component: AboutView
  },
  { 
    path: '/catalog/:id', 
    name: 'card',
    component: CardView 
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
