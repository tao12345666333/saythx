<template>
  <div class="say-thanks">
    <el-row :gutter="20" v-loading.fullscreen.lock="loading">
      <el-col :span="16">
        <el-form ref="form" :model="honorData" label-width="80px">
          <el-form-item label="Thanks to">
            <el-input v-model="honorData.to"></el-input>
          </el-form-item>
          <el-form-item label="Thank you">
            <el-input type="textarea" v-model="honorData.content" rows="15"></el-input>
          </el-form-item>
          <el-form-item label="From">
            <el-input v-model="honorData.from"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">Send Note</el-button>
          </el-form-item>
        </el-form>
      </el-col>
      <el-col :span="8">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>HonorList</span>
            <el-button style="float: right; padding: 3px 0" type="text" @click="getList">Refresh</el-button>
          </div>
          <el-table
            :data="honorList"
            stripe
            style="width: 100%">
            <el-table-column
              prop="name"
              label="Name"
              width="180">
            </el-table-column>
            <el-table-column
              prop="score"
              label="Score"
              width="180">
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: 'SayThanks',
  data() {
    return {
      honorData: {
        to: '',
        from: '',
        content: '',
      },
      loading: false,
      honorList: [],
    }
  },
  methods: {
    onSubmit() {
      this.loading = true;
      this.$axios.post('/api/v1/saythx', this.honorData).then(
        () => {
          this.loading = false;
          this.honorData = {
            to: '',
            from: '',
            content: '',
          }
          this.getList()
          this.$message({
            message: 'Thanks successed',
            type: 'success'
          });
      }).catch(
        (response) => {
          this.$message({
            message: 'Oops!',
            type: 'warning'
          });
        }
      )
    },
    getList() {
      this.$axios.get('/api/v1/list').then((response) => (
        this.honorList = response.data.HonorDatas
      ))
    },

  },
  mounted() {
    this.getList()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
