import axios from "axios";

export const V2rayConfig = {
    get: function () {
        return axios.get("/controller/v2ray/config").catch(error => {
            throw new Error(`[V2rayRawConfig] get ${error}`);
        });
    },
    update: function (data) {
        return axios.put("/controller/v2ray/config", data).catch(error => {
            throw new Error(`[V2rayRawConfig] update ${error}`);
        })
    }
}

export const Nodes = {
    queryAll: function () {
        return axios.get("/api/nodes").catch(error => {
            throw new Error(`[Nodes] query ${error}`)
        });
    },
    create: function (data) {
        return axios.post("/api/nodes", data).catch(error => {
            throw new Error(`[Nodes] create ${error}`)
        });
    },
    remove: function (name) {
        return axios.delete("/api/nodes/" + name).catch(error => {
            throw new Error(`[Nodes] remove ${error}`)
        });
    },
    update: function (name, data) {
        axios.put("/api/nodes/" + name, data).catch(error => {
            throw new Error(`[Nodes] update ${error}`)
        });
    }
}