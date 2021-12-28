<template>
  <el-container style="height: 500px; border: 1px solid #eee">
    <el-aside width="200px" style="background-color: rgb(238, 241, 246)">
      <el-menu
          :unique-opened="true"
          :router="true"
          @open="handleOpen"
          @close="handleClose"
          default-active="/"
          class="el-menu-vertical-demo"
          background-color="#545c64"
          text-color="#fff"
          active-text-color="#ffd04b">
        <el-menu-item index="/">
          <i class="el-icon-s-home"></i>Smart Home Gateway
        </el-menu-item>

        <el-menu-item index="/status">
          <i class="el-icon-star-on"></i>Status
        </el-menu-item>
        <el-submenu index="/v2ray">
          <template #title><i class="el-icon-s-promotion"></i>V2ray</template>
          <el-menu-item-group>
            <el-menu-item index="/v2ray-node">Node</el-menu-item>
          </el-menu-item-group>
          <el-menu-item-group>
            <el-menu-item index="/v2ray-setting">Setting</el-menu-item>
          </el-menu-item-group>
        </el-submenu>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header style="text-align: right; font-size: 18px">
        <el-dropdown>
          <i class="el-icon-user-solid" style="margin-right: 15px"></i>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item>login out</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
        <span>admin</span>
      </el-header>
      <el-main>
        <router-view/>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
export default {
  data() {
    return {}
  },
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    }
  }
};
</script>
<style>
.el-header {
  background-color: #dbe1ea;
  color: #333;
  line-height: 60px;
}

.el-main {
  height: 500px;
}
</style>