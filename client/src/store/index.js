import Vue from "vue";
import Vuex from "vuex";
import config from "@/store/config";

Vue.use(Vuex);

export default new Vuex.Store({
    modules: {
        config
    }
})