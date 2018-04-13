import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import About from '@/components/About'
import Index from '@/components/Index'
import LoginForm from '@/components/LoginForm'

Vue.use(Router)

const routes = [
	{
		path: '/',
		name: 'Index',
		component: Index
	}, 
	{
		path: '/home',
		name: 'Home',
		component: Home
	},
	{
		path: '/about',
		name: 'About',
		component: About
	},
	{
		path: '/login',
		name: 'login',
		component: LoginForm
	}
]

const router = new Router({
	mode: 'history',
	routes: routes
})

export default router