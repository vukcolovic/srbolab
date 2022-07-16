import { createRouter, createWebHistory } from 'vue-router'
import BodyComponent from './../components/BodyComponent.vue'
import LoginComponent from './../components/LoginComponent.vue'
import UsersList from '../components/UsersList.vue'
import UserEdit from './../components/UserEdit.vue'
import SupportComponent from './../components/SupportComponent.vue'
import IrregularityLevels from './../components/IrregularityLevels.vue'
import IrregularitiesComponent from '../components/IrregularitiesList.vue'
import IrregularityEdit from './../components/IrregularityEdit.vue'
import FuelConsumptionList from './../components/FuelConsumptionList.vue'
import FuelConsumptionEdit from './../components/FuelConsumptionEdit.vue'


const routes = [
    {
        path: '/',
        name: 'Home',
        component: BodyComponent,
    },
    {
        path: '/login',
        name: 'Login',
        component: LoginComponent,
    },
    {
        path: '/users',
        name: 'Users',
        component: UsersList,
    },
    {
        path: '/user',
        name: 'UserEdit',
        component: UserEdit,
        props: true,
    },
    {
        path: '/support',
        name: 'SupportComponent',
        component: SupportComponent,
    },
    {
        path: '/irregularity-levels',
        name: 'IrregularityLevels',
        component: IrregularityLevels,
    },
    {
    path: '/irregularities',
    name: 'IrregularitiesComponent',
    component: IrregularitiesComponent,
    },
    {
        path: '/irregularity',
        name: 'IrregularityEdit',
        component: IrregularityEdit,
        props: true,
    },
    {
        path: '/fuel',
        name: 'FuelConsumptionList',
        component: FuelConsumptionList,
        props: true,
    },
    {
        path: '/fuel-edit',
        name: 'FuelConsumptionEdit',
        component: FuelConsumptionEdit,
        props: true,
    }
]

const router = createRouter({history: createWebHistory(), routes})
export default router
