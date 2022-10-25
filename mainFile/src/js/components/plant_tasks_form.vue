<template>
  <div>
    <el-form ref="form" :model="formData" :rules="rules" :label-width="formLabelWidth">

      <el-row>
        <el-col :span="12" :xs="24">
          <el-form-item label="任务名称">
            <el-input v-model="formData.name">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12" :xs="24">
          <el-form-item label="执行单位">
            <el-input v-model="formData.organization">
            </el-input>
          </el-form-item>
        </el-col>
      </el-row>


      <el-form-item label="任务内容" required>
        <template v-if="formDataTaskSettingBuild">
          <div v-for="(ts,tsi) in $LeUtils.valueGet(formData,'task_setting',[])">
            <el-descriptions class="" title="" :column="8" size="small" :colon="false">
              <el-descriptions-item label="任务类型" :span=2>
                <el-select size="small" v-model="formData['task_setting'][tsi]['n']">
                  <el-option
                      key="ip"
                      label="IP扫描"
                      value="ip">
                  </el-option>
                  <el-option
                      key="port"
                      label="端口扫描"
                      value="port">
                  </el-option>
                  <el-option
                      key="blast"
                      label="口令爆破"
                      value="blast">
                  </el-option>
                  <el-option
                      key="bug"
                      label="漏洞探测"
                      value="bug">
                  </el-option>
                  <el-option
                      key="file"
                      label="文件遍历"
                      value="file">
                  </el-option>
                </el-select>
              </el-descriptions-item>
              <el-descriptions-item label="时长">
                <el-input size="small" type="number" v-model.number="formData['task_setting'][tsi]['s']">
                  <template v-slot:suffix>
                    <div style="height:30px;line-height: 30px">秒</div>
                  </template>
                </el-input>
              </el-descriptions-item>
              <el-descriptions-item label="频率">
                <el-input size="small" type="number" v-model.number="formData['task_setting'][tsi]['t']">
                </el-input>
              </el-descriptions-item>
              <el-descriptions-item label="并发">
                <el-input size="small" type="number" v-model.number="formData['task_setting'][tsi]['mt']">
                </el-input>
              </el-descriptions-item>
              <el-descriptions-item label="端口" :span="2">
                <el-checkbox-group v-model="formData['task_setting'][tsi]['port']">
                  <el-checkbox :label=3389>3389</el-checkbox>
                  <el-checkbox :label=22>22</el-checkbox>
                  <el-checkbox :label=445>445</el-checkbox>
                </el-checkbox-group>
              </el-descriptions-item>
              <el-descriptions-item label="">
                <el-button size="mini" type="danger" plain @click="taskSettingRemoveItem(tsi)">移除内容</el-button>
              </el-descriptions-item>
            </el-descriptions>
          </div>
        </template>
        <div>
          <el-button size="mini" type="primary" @click="taskSettingAddItem()">添加内容</el-button>
        </div>
      </el-form-item>


      <el-row>
        <el-col :span="6" :xs="12">
          <el-form-item label="植入时间" required>

            <el-date-picker
                v-model="formData.implant_time"
                type="datetime"
                value-format="yyyy-MM-dd HH:mm:ss"
                placeholder="选择日期时间">
            </el-date-picker>
          </el-form-item>
        </el-col>
        <el-col :span="6" :xs="12">
          <el-form-item label="激活时间" required>
            <el-date-picker
                v-model="formData.activated_time"
                type="datetime"
                value-format="yyyy-MM-dd HH:mm:ss"
                placeholder="选择日期时间">
            </el-date-picker>

          </el-form-item>
        </el-col>
        <el-col :span="6" :xs="12">
          <el-form-item label="停止时间" required>
            <el-date-picker
                v-model="formData.end_time"
                type="datetime"
                value-format="yyyy-MM-dd HH:mm:ss"
                placeholder="选择日期时间">
            </el-date-picker>
          </el-form-item>
        </el-col>
        <el-col :span="6" :xs="12">
          <el-form-item label="自毁时间" required>
            <el-date-picker
                v-model="formData.self_delete_time"
                type="datetime"
                value-format="yyyy-MM-dd HH:mm:ss"
                placeholder="选择日期时间">
            </el-date-picker>
          </el-form-item>
        </el-col>
      </el-row>

      <el-row>
        <el-col :span="12" :xs="24">
          <el-form-item label="回连地址" required>
            <el-input v-model="formData.service_ip">
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

    dataEntity: {
      type: Object,
      default: null
    },
    statusOptions: {
      type: Array,
      default: function () {
        return [
          {
            "label": "是",
            "value": "1",
          },
          {
            "label": "否",
            "value": "0",
          },
        ]
      }
    },

  },
  data() {

    return {
      formLabelWidth: window.formLabelWidth || '120px',
      formData: {
        task_setting: []
      },
      formDataTaskSettingBuild: true,
      formErrors: {},
      rules: {

      },


    }
  },
  created() {
    this.assignData(this.dataEntity);
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
      /*for(var i in LeUtils.valueGet(this.formData,'task_setting',[])){
        var item = LeUtils.valueGet(this.formData,'task_setting.'+i,[])
        console.log(item)
        this.formDataTaskSetting.push(item)
      }*/




    },

    submit: function () {
      let _this = this;

      this.$refs['form'].validate((valid) => {
        if (valid) {

          LeRoute.request({
            uri: "/console/plant-tasks/entity/{id}",
            method: "post"
          }, _this.formData).then(function (res_data) {
            let id = LeUtils.valueGet(res_data, 'data.id')
            _this.$message.success({
              message: '保存成功',
              onClose: function () {
                if(id){
                  window.location.href="/console/plant-tasks/entity/"+id
                }else{
                  window.location.reload()
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
