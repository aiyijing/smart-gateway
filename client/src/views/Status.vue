<template>
  <el-row>
    <el-col :span="4">
      <el-card :body-style="{ padding: '0px' }">
        <div>
          <el-progress type="circle" :percentage="100" status="success"></el-progress>
        </div>
        <div style="padding: 14px;">
          <span>V2ray Service</span>
          <div class="bottom">
            <time class="time">Running 24h</time>
            <el-button type="primary" class="button" v-on:click="restart">Restart</el-button>
          </div>
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>


<script>
export default {
  data() {
    return {}
  },
  methods: {
    failedCallBack: function () {
      const h = this.$createElement;
      this.$notify({
        title: 'Notification',
        message: h('i', {style: 'color: teal'}, 'Restart Failed!')
      });
    },
    successCallBack: function () {
      const h = this.$createElement;
      this.$notify({
        title: 'Notification',
        message: h('i', {style: 'color: teal'}, 'Restart Success!')
      });
    },
    restart: function () {
      let action_endpoint = "/controller/v2ray/action?action=restart";
      let status = this
      let xhr = new XMLHttpRequest();
      xhr.open("GET", action_endpoint, true)
      xhr.send();

      xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
          if (xhr.status === 200) {
            status.successCallBack();
          } else {
            status.failedCallBack();
          }
        }
      };
    }
  }
}
</script>


<style>

.time {
  font-size: 13px;
  color: #999;
}

.bottom {
  margin-top: 13px;
  line-height: 12px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.button {
  padding: 5px;
}

</style>