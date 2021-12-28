<template>
  <div>
    <el-row>
      <el-button type="primary" @click="dialogFormVisible = true">添加节点</el-button>

      <el-dialog title="添加节点" :visible.sync="dialogFormVisible">
        <el-form :model="form">
          <el-form-item label="名称" :label-width="'120px'">
            <el-input v-model="form.name"></el-input>
          </el-form-item>
          <el-form-item label="OutBound" :label-width="'120px'">
            <el-input
                type="textarea"
                :rows="8"
                placeholder="outbound"
                v-model="form.outbound">
            </el-input>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="dialogFormVisible = false">取 消</el-button>
          <el-button type="primary" @click="add">确 定</el-button>
        </div>
      </el-dialog>
    </el-row>
    <el-row>
      <el-table :data="tableData" style="width: 100%">
        <el-table-column prop="name" label="节点名称" width="180">
        </el-table-column>
        <el-table-column prop="address" label="地址"></el-table-column>
        <el-table-column prop="enable" label="状态" width="180">
          <template slot-scope="scope">
            <el-switch v-model="scope.row.enable"
                       @change="updateNode(scope.row.name,scope.row.enable)"
                       active-color="#13ce66"
                       inactive-color="#ff4949"></el-switch>
          </template>
        </el-table-column>
        <el-table-column>
          <template slot-scope="scope">
            <el-button v-on:click="edit(scope.row.name)" size="mini">edit</el-button>
            <el-button v-on:click="del(scope.row.name)" size="mini" type="danger">delete</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-row>
  </div>
</template>

<script>
import {Nodes} from "@/api"

export default {
  data() {
    return {
      nodes: [],
      status: true,
      dialogFormVisible: false,
      form: {
        name: '',
        outbound: '',
        enable: false,
      }
    }
  },
  mounted() {
    this.queryAll()
  },
  computed: {
    tableData: function () {
      return this.nodes.map(function (node) {
        let address = ""
        try {
          address = node.outBound.settings.vnext[0].address
        } catch (e) {
          address = ""
        }
        return {
          name: node.name,
          enable: node.enable,
          address: address,
        }
      })
    }
  },
  methods: {
    queryAll: async function () {
      let {data, status} = await Nodes.queryAll()
      if (status !== 200) {
        this.notify("failed to query outBound")
        return
      }
      this.nodes = data.data
    },
    async updateNode(name, enable) {
      let node = this.getNodeByName(name)
      node.enable = enable
      let {status} = await Nodes.update(name, node)
      if (status === 200) {
        this.notify("switch success!")
      } else {
        this.notify("switch failed!")
        node.enable = !enable
      }
    },
    edit: function (name) {
      console.log("update:" + name)
    },
    del: async function (name) {
      let {status} = await Nodes.remove(name)
      if (status === 200) {
        this.notify("delete success")
        this.removeNodeByName(name)
      } else {
        this.notify("delete failed")
      }
    },
    add: async function () {
      if (this.form.name === "") {
        this.notify("name not allow empty!")
        return
      }
      if (this.form.outbound === "") {
        this.notify("outbound not allow empty!")
        return
      }

      let outbound = {}
      try {
        outbound = JSON.parse(this.form.outbound)
      } catch (e) {
        this.notify("create failed" + e)
        return
      }
      let node = {
        name: this.form.name,
        enable: this.form.enable,
        outbound: outbound
      }
      let {data, status} = await Nodes.create(node)
      if (status === 200) {
        this.notify("create success")
      } else {
        this.notify("create failed")
      }
      this.nodes.push(data.data)

      this.dialogFormVisible = false
      this.form.outbound = ""
      this.form.name = ""
    },
    notify(msg) {
      const h = this.$createElement;
      this.$notify({
        title: 'Notification',
        message: h('i', {style: 'color: teal'}, msg)
      });
    },
    removeNodeByName(name) {
      for (let i = 0; this.nodes.length; i++) {
        if (this.nodes[i].name === name) {
          this.nodes.splice(i, 1);
        }
      }
    },
    getNodeByName(name) {
      for (let node of this.nodes) {
        if (node.name === name) {
          return node
        }
      }
    }
  }
}
</script>