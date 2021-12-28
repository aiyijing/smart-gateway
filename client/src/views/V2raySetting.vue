<template>
  <div>
    <el-row>
      <el-button type="primary" v-on:click="fetchConfigConfig">Refresh</el-button>
    </el-row>
    <el-row>
      <el-input
          type="textarea"
          :readonly="true"
          :rows="16"
          placeholder="V2ray config.json"
          v-model="config">
      </el-input>
    </el-row>
  </div>
</template>

<script>

import {V2rayConfig} from "@/api"

export default {
  data() {
    return {
      config: ""
    }
  },
  mounted() {
    this.fetchConfigConfig()
  },
  methods: {
    async fetchConfigConfig() {
      let {data, status} = await V2rayConfig.get()
      if (status === 200) {
        this.notify("refresh success!")
      } else {
        this.notify("refresh failed!")
      }
      this.config = JSON.stringify(data, null, 2);
    },
    notify(msg) {
      const h = this.$createElement;
      this.$notify({
        title: 'Notification',
        message: h('i', {style: 'color: teal'}, msg)
      });
    },
    alert: function (msg) {
      this.$alert(msg, 'Message', {
        confirmButtonText: '确定',
      });
    }
  }
}
</script>
<style>
.el-row {
  margin-bottom: 20px;

}

:last-child {
  margin-bottom: 0;
}

</style>