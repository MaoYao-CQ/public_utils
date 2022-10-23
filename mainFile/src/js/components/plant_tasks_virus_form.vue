
<template>
  <div >
    <el-form ref="form" :model="formData" :rules="rules" :label-width="formLabelWidth">

      <el-row>
        <el-col :span="8" :xs="24">
          <el-form-item label="选择病毒" required>
            <form-input-select-virus
                :data-only-unused="true"
                :data-value="formData.id"
                :disabled="disableChangeId"
                v-on:change="(val)=>{formData.id=val}"
            >

            </form-input-select-virus>
          </el-form-item>
        </el-col>
        <el-col :span="8" :xs="24">
          <el-form-item label="日志路径" required>
            <el-input v-model="formData.log_path">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="8" :xs="24">
          <el-form-item label="执行人" >
            <el-input v-model="formData.executor">
            </el-input>
          </el-form-item>
        </el-col>

      </el-row>

      <el-row>
        <el-col :span="8" :xs="24">
          <el-form-item label="网络区域" >
            <el-input v-model="formData.network_area">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="8" :xs="24">
          <el-form-item label="目标IP" >
            <el-input v-model="formData.target_ip">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="8" :xs="24">
          <el-form-item label="网络MAC" >
            <el-input v-model="formData.target_mac">
            </el-input>
          </el-form-item>
        </el-col>

      </el-row>

      
      <el-form-item label="">
        <el-button type="primary" size="small" @click="submit">保存</el-button>

      </el-form-item>
    </el-form>


  </div>

</template>

<script>
export default {
  components: {},
  props: {
    dataTask: {
      type: Object,
      default: function (){
        return null
      }
    },
    dataEntity: {
      type: Object,
      default: function (){
        return null
      }
    },


  },
  data() {

    return {
      formLabelWidth: window.formLabelWidth || '120px',
      formData: {
        task_id:"",
        id:"",
        target_ip:"",
        target_mac:"",
        network_area:"",
        executor:"",
        log_path:"",
      },
      disableChangeId : false,
      formErrors: {},
      rules: {

      },


    }
  },
  created() {
    this.assignData();
  },
  mounted() {
    let _this = this;
  },
  methods: {
    assignData: function () {
      if (!!this.dataEntity && LeUtils.typeIs('object', this.dataEntity) && !LeUtils.isEmpty(this.dataEntity)) {
        this.formData = Object.assign({}, this.formData, this.dataEntity)
        this.disableChangeId = true
      }

      if (!!this.dataTask && LeUtils.typeIs('object', this.dataTask)) {
        this.formData.task_id = LeUtils.valueGet(this.dataTask,'id',"")
      }


    },

    submit: function () {
      let _this = this;

      LeRoute.request({
        uri: "/console/virus/assign-task/{id}",
        method: "post"
      }, _this.formData).then(function (res_data) {
        let id = LeUtils.valueGet(res_data, 'data.id')
        _this.$message.success({
          message: '保存成功',
          onClose: function () {
            window.location.reload()
          }
        })


      })
    },


  }
}
</script>
