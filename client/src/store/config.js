import { V2rayConfig } from "@/api";

const configState = {
    config: "aiyijing"
}

const state = {
    ...configState
}

const mutations = {
    update(state, config) {
        state.config = config;
    }
};

const actions = {
    async fetchConfig(context) {
        const {data, status} = await V2rayConfig.get();
        if (status === 200) {
            await context.commit("update", JSON.stringify(data, null, 2));
        }
    },
    async updateConfig(context, newConfig) {
        const {status} = await V2rayConfig.update(newConfig);
        if (status === 200) {
            context.commit("update", JSON.stringify(newConfig, null, 2))
        }
    }
}

const getters = {
    config(state) {
        return state.config;
    },
}

export default {
    mutations,
    actions,
    state,
    getters,
}