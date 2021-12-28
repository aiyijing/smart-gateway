import VueRouter from 'vue-router';
import Vue from 'vue'
import Status from '../views/Status.vue';
import V2raySetting from '../views/V2raySetting.vue';
import Node from '../views/Node.vue';

Vue.use(VueRouter)

const router = new VueRouter({
    routes: [
        {
            path: '/status',
            component: Status,
        },
        {
            path: '/v2ray-node',
            component: Node
        },
        {
            path: '/v2ray-setting',
            component: V2raySetting
        }
    ]
})

export default router