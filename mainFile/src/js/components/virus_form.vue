<template>
  <div>
    <el-form ref="form" :model="formData" :rules="rules" :label-width="formLabelWidth">

      <el-row v-if="!isAddNew">
        <el-col :span="12" :xs="24">
          <el-form-item label="病毒ID">
            {{dataEntity.id}}
          </el-form-item>
        </el-col>
        <el-col :span="12" :xs="24">

        </el-col>
      </el-row>

      <el-row>
        <el-col :span="12" :xs="24">
          <el-form-item label="目标系统" required>

            <el-select v-model="formData.run_os"  :disabled="!isAddNew" >
              <el-option
                  label="Windows"
                  value="windows">
              </el-option>
              <el-option
                  label="Linux"
                  value="linux">
              </el-option>
            </el-select>

          </el-form-item>
        </el-col>
        <el-col :span="12" :xs="24">
<!--
          <el-form-item label="目标系统架构">
            <el-select v-model="formData.target_os_arch" :disabled="!isAddNew" >
              <el-option
                  label="Intel x64/AMD64"
                  value="amd64">
              </el-option>
              <el-option
                  label="Intel x86"
                  value="386">
              </el-option>
            </el-select>
          </el-form-item>
        -->
        </el-col>
      </el-row>


      <el-row>
        <el-col :span="12" :xs="24">
          <el-form-item label="病毒名称">
            <el-input v-model="formData.name">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12" :xs="24">

        </el-col>

      </el-row>
      <el-row>
        <el-col :span="12" :xs="24">
          <el-form-item label="生成数量" v-if="isAddNew">
            <el-input type="number" v-model.number="formData.quantity">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12" :xs="24">

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

    dataEntity: {
      type: Object,
      default: null
    },
  },
  data() {

    return {
      formLabelWidth: window.formLabelWidth || '120px',
      formData: {

      },
      formErrors: {},
      rules: {

      },
    }
  },
  computed: {
    isAddNew: function () {
      return !LeUtils.valueGet(this.dataEntity, 'id','');
    }
  },
  created() {
    this.assignData(this.dataEntity);
    //console.log(this.dataEntity)
  },
  mounted() {
    let _this = this;
    //console.log(_this.formData)
  },
  methods: {
    assignData: function (data) {
      if (LeUtils.typeIs('object', data)) {
        this.formData = Object.assign({}, this.formData, data)

      }

    },

    submit: function () {
      let _this = this;

      this.$refs['form'].validate((valid) => {
        if (valid) {

          LeRoute.request({
            uri: "/console/virus/entity/{id}",
            method: "post"
          }, _this.formData).then(function (res_data) {
            let id = LeUtils.valueGet(res_data, 'data.id')
            _this.$message.success({
              message: '保存成功',
              onClose: function () {
                if(id){
                  window.location.href="/console/virus/entity/"+id
                }else{
                  window.location.href = "/console/virus/index"
                }

              }
            })


          })

        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },


    taskSettingAddItem: function () {
      this.formData.task_setting.push({n:"",s:"",t:"",mt:"",port:[]})

    },
    taskSettingRemoveItem: function (tsi) {
      //let _this = this
      //this.formDataTaskSettingBuild = false
      this.formData.task_setting.splice(tsi, 1)
      /*this.$nextTick(function () {
        console.log(_this.formData.task_setting)
        _this.formDataTaskSettingBuild = true
      })*/
    }
  }
}
</script>
