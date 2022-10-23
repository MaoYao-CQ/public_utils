<template>
    <el-select v-model="selectValue" filterable v-on:change="handleChange" remote :remote-method="loadOptions" :disabled="disabled" >

        <el-option
            v-for="item in selectOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value">
        </el-option>
    </el-select>

</template>

<script>

export default {
    props: {
        dataValue: {
            type: String,Number,
            default: function () {
                return ''
            }
        },
        dataOnlyUnused: {
            type: Boolean,
            default: true
        },
        disabled: {
            type: Boolean,
            default: true
        },
    },
    data() {
        return {
            selectOptions: [],
            selectValue: '',

        }
    },
    created() {
        let _this = this;
        _this.selectValue = _this.dataValue
        _this.loadOptions();
    },
    mounted() {

    },
    methods: {
        handleChange: function (val) {
            this.$emit('change',val)
        },
        loadOptions:function (query){
            let _this = this;
            LeRoute.quiet().request({uri: "/console/virus/select-options", method: "get"}, {
                selected: _this.selectValue,
              only_unused: _this.dataOnlyUnused ,
            }).then(function (res_data) {
                _this.selectOptions = LeUtils.valueGet(res_data,'data',[])
            })
        }

    }
}
</script>
