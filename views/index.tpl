{{template "layout_header.tpl" .}}
<style>
    .self-cord {
        height: 395px;
        min-width: 490px;
        overflow-y: scroll;
    }

    .self-cord .string-box-p {
        margin-top: 15px;
    }

    .container .row {
        height: 48px;
    }
</style>
<div class="container w-100 h-100 d-block mt-3 main-content">
    <div class="d-block">
        <div class="self-cord w-50 d-inline-block p-1 float-left overflow-auto">
            <div class="card text-white bg-secondary h-100">
                <div class="card-header">随机字符串</div>
                <div class="card-body" id="random-string-box-1">
                    <div class="w-100 text-center">
                        <div class="form-check form-check-inline">
                            <input class="form-check-input string-contain-number" type="checkbox"
                                   id="inlineCheckbox-1-1"
                                   value="number"
                                   onchange="randomStringBox1()"
                                   checked>
                            <label class="form-check-label" for="inlineCheckbox-1-1">数字</label>
                        </div>
                        <div class="form-check form-check-inline">
                            <input class="form-check-input string-contain-uppercase" type="checkbox"
                                   id="inlineCheckbox-1-2"
                                   value="uppercase"
                                   onchange="randomStringBox1()"
                                   checked>
                            <label class="form-check-label" for="inlineCheckbox-1-2">大写</label>
                        </div>
                        <div class="form-check form-check-inline">
                            <input class="form-check-input string-contain-lowercase" type="checkbox"
                                   id="inlineCheckbox-1-3"
                                   value="lowercase"
                                   onchange="randomStringBox1()"
                                   checked>
                            <label class="form-check-label" for="inlineCheckbox-1-3">小写</label>
                        </div>
                        <div class="form-check form-check-inline">
                            <input class="form-check-input string-contain-special" type="checkbox"
                                   id="inlineCheckbox-1-4"
                                   value="special"
                                   onchange="randomStringBox1()"
                            >
                            <label class="form-check-label" for="inlineCheckbox-1-4">特殊</label>
                        </div>
                        <div class="form-check form-check-inline">
                            <input class="form-check-input string-contain-take-out-confused" type="checkbox"
                                   id="inlineCheckbox-1-5"
                                   onchange="randomStringBox1()"
                                   value="take_out_confused" checked>
                            <label class="form-check-label" for="inlineCheckbox-1-5">去掉易混淆</label>
                        </div>
                        <button type="button" class="btn btn-primary btn-sm" onclick="randomStringBox1()">换一批</button>
                    </div>
                    <p class="card-text text-center overflow-auto string-box-p"></p>
                </div>
            </div>
        </div>
        <div class="self-cord w-50 d-inline-block p-1 float-left overflow-auto">
            <div class="card text-white bg-secondary h-100">
                <div class="card-header">时间、IP</div>
                <div class="w-100 text-center">

                </div>
                <div class="container mt-3">
                    <div class="row align-items-center">
                        <div class="col-3 text-right">
                            时间
                        </div>
                        <div class="col-auto">
                            <span id="timestamp">{{.timestamp}}</span><span class="ml-5" id="datetime"></span>
                        </div>
                    </div>

                    <div class="row align-items-center">
                        <div class="col-3 text-right">
                            unix
                        </div>
                        <div class="col-auto">
                            <div class="input-group input-group-sm float-left" style="width: 190px">
                                <input type="text" class="form-control unixToDatetime-input" placeholder="请输入时间戳"
                                       aria-label="请输入时间戳">
                                <div class="input-group-append">
                                    <button class="btn btn-primary" type="button" id="unixToDatetime">转换</button>
                                </div>
                            </div>
                            <span class="ml-2 d-block float-left unixToDatetime-show"></span>
                            <div class="clearfix"></div>
                        </div>
                    </div>
                    <div class="row align-items-center">
                        <div class="col-3 text-right">
                            datetime
                        </div>
                        <div class="col-auto">
                            <div class="input-group input-group-sm float-left" style="width: 190px">
                                <input type="text" class="form-control datetimeToUnix-input" placeholder="请输入时间"
                                       aria-label="请输入时间" value="{{.datetime}}">
                                <div class="input-group-append">
                                    <button class="btn btn-primary" type="button" id="datetimeToUnix">转换</button>
                                </div>
                            </div>
                            <span class="ml-2 d-block float-left datetimeToUnix-show"></span>
                            <div class="clearfix"></div>
                        </div>
                    </div>
                    <div class="row mt-2">
                        <div class="col-3 text-right">
                            IP查询
                        </div>
                        <div class="col-auto">
                            <div class="input-group input-group-sm float-left" style="width: 190px">
                                <input type="text" class="form-control ipQuery-input" placeholder="请输入IP地址"
                                       aria-label="请输入IP地址" value="">
                                <div class="input-group-append">
                                    <button class="btn btn-primary" type="button" id="ipQuery">查询</button>
                                </div>
                            </div>
                            <span class="ml-2 d-block float-left">默认查询自己的IP</span>
                            <div class="clearfix"></div>
                            <span class="ml-2 d-block ipQuery-show"></span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="clearfix"></div>
    <div class="h-50 d-block">
        <div class="self-cord w-50 d-inline-block p-1 float-left overflow-auto">
            <div class="card text-white bg-secondary h-100">
                <div class="card-header">随机字符串</div>
                <div class="card-body" id="random-string-box-2">
                    <div class="w-100 text-center">
                        <div class="form-check form-check-inline">
                            <input class="form-check-input string-contain-number" type="checkbox"
                                   id="inlineCheckbox-2-1"
                                   value="number"
                                   onchange="randomStringBox2()"
                                   checked>
                            <label class="form-check-label" for="inlineCheckbox-2-1">数字</label>
                        </div>
                        <div class="form-check form-check-inline">
                            <input class="form-check-input string-contain-uppercase" type="checkbox"
                                   id="inlineCheckbox-2-2"
                                   value="uppercase"
                                   onchange="randomStringBox2()">
                            <label class="form-check-label" for="inlineCheckbox-2-2">大写</label>
                        </div>
                        <div class="form-check form-check-inline">
                            <input class="form-check-input string-contain-lowercase" type="checkbox"
                                   id="inlineCheckbox-2-3"
                                   value="lowercase"
                                   onchange="randomStringBox2()"
                                   checked>
                            <label class="form-check-label" for="inlineCheckbox-2-3">小写</label>
                        </div>
                        <div class="form-check form-check-inline">
                            <input class="form-check-input string-contain-special" type="checkbox"
                                   id="inlineCheckbox-2-4"
                                   value="special"
                                   onchange="randomStringBox2()"
                            >
                            <label class="form-check-label" for="inlineCheckbox-2-4">特殊</label>
                        </div>
                        <div class="form-check form-check-inline">
                            <input class="form-check-input string-contain-take-out-confused" type="checkbox"
                                   id="inlineCheckbox-2-5"
                                   onchange="randomStringBox2()"
                                   value="take_out_confused" checked>
                            <label class="form-check-label" for="inlineCheckbox-2-5">去掉易混淆</label>
                        </div>
                        <button type="button" class="btn btn-primary btn-sm" onclick="randomStringBox2()">换一批</button>
                    </div>
                    <p class="card-text text-center overflow-auto string-box-p"></p>
                </div>
            </div>
        </div>
        <div class="self-cord w-50 d-inline-block p-1 float-left overflow-auto">
            <div class="card text-white bg-secondary h-100">
                <div class="card-header">随机字符串</div>
                <div class="card-body" id="random-string-box-3">
                    <div class="w-100 text-center">
                        <div class="form-check form-check-inline">
                            <input class="form-check-input string-contain-number" type="checkbox"
                                   id="inlineCheckbox-3-1"
                                   value="number"
                                   onchange="randomStringBox3()"
                                   checked>
                            <label class="form-check-label" for="inlineCheckbox-3-1">数字</label>
                        </div>
                        <div class="form-check form-check-inline">
                            <input class="form-check-input string-contain-uppercase" type="checkbox"
                                   id="inlineCheckbox-3-2"
                                   value="uppercase"
                                   onchange="randomStringBox3()" checked>
                            <label class="form-check-label" for="inlineCheckbox-3-2">大写</label>
                        </div>
                        <div class="form-check form-check-inline">
                            <input class="form-check-input string-contain-lowercase" type="checkbox"
                                   id="inlineCheckbox-3-3"
                                   value="lowercase"
                                   onchange="randomStringBox3()"
                                   checked>
                            <label class="form-check-label" for="inlineCheckbox-3-3">小写</label>
                        </div>
                        <div class="form-check form-check-inline">
                            <input class="form-check-input string-contain-special" type="checkbox"
                                   id="inlineCheckbox-3-4"
                                   value="special"
                                   onchange="randomStringBox3()" checked
                            >
                            <label class="form-check-label" for="inlineCheckbox-3-4">特殊</label>
                        </div>
                        <div class="form-check form-check-inline">
                            <input class="form-check-input string-contain-take-out-confused" type="checkbox"
                                   id="inlineCheckbox-3-5"
                                   onchange="randomStringBox3()"
                                   value="take_out_confused" checked>
                            <label class="form-check-label" for="inlineCheckbox-3-5">去掉易混淆</label>
                        </div>
                        <button type="button" class="btn btn-primary btn-sm" onclick="randomStringBox3()">换一批</button>
                    </div>
                    <p class="card-text text-center overflow-auto string-box-p"></p>
                </div>
            </div>
        </div>
    </div>
</div>
<div class="clearfix"></div>

<script>
    // let timestamp = $("#timestamp").text();
    let timestamp = "{{.timestamp}}";
    let _timestamp = timestamp * 1000;

    Date.prototype.format = function (fmt) {
        var o = {
            "M+": this.getMonth() + 1, //月份
            "d+": this.getDate(), //日
            "h+": this.getHours(), //小时
            "m+": this.getMinutes(), //分
            "s+": this.getSeconds(), //秒
            "q+": Math.floor((this.getMonth() + 3) / 3), //季度
            "S": this.getMilliseconds() //毫秒
        };
        if (/(y+)/.test(fmt)) {
            fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
        }
        for (var k in o) {
            if (new RegExp("(" + k + ")").test(fmt)) {
                fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
            }
        }
        return fmt;
    }

    function getTime() {
        _timestamp += 1000;
        let t = new Date(_timestamp)
        $("#timestamp").text(t.getTime() / 1000)
        $("#datetime").text(t.format("yyyy年MM月dd日 hh:mm:ss"))
    }

    setInterval(getTime, 1000);

    $('#unixToDatetime').click(function () {
        let timestamp = $('.unixToDatetime-input').val();

        if ((timestamp.length != 10 && timestamp.length != 13) || timestamp != Number(timestamp)) {
            $(".unixToDatetime-show").html("NaN");
            return
        }
        if (timestamp.length == 10) {
            timestamp *= 1000;
        } else {
            timestamp *= 1;
        }

        $(".unixToDatetime-show").html(new Date(timestamp).format("yyyy-MM-dd hh:mm:ss"))
    });


    $('#datetimeToUnix').click(function () {
        let datetime = $('.datetimeToUnix-input').val();

        let t = new Date(datetime)

        $(".datetimeToUnix-show").html(t.getTime() / 1000)
    });

    $('#ipQuery').click(function () {
        let ip = $('.ipQuery-input').val()
        if (ip === "") ip = "self"
        $.ajax({
            method:"get",
            data:{ip:ip},
            url: "/api/v1/ip/get",
            success:function (ret) {
                if (ret.status){
                    data = ret.data;
                    $('.ipQuery-input').val(data.ip);
                    let s = "归属："+data.country_name + " " + data.region_name + " " + data.city_name;
                    console.log(s)
                    $('.ipQuery-show').html(s);
                } else {
                    console.log(ret.error)
                    alert(ret.message)
                }
            }
        })
    });


    $('#collect').click(function () {
        let url = 'http://www.ai996.top'
        let name = '工具喵'
        try {
            window.external.AddFavorite(url, name)
        } catch (e1) {
            try {
                window.sidebar.addPanel(name, url, "");
            } catch (e2) {
                // alert("加入收藏失败，请使用Ctrl D进行添加"+e1.message+" e2:"+e2.message)
                alert("加入收藏失败，请使用Ctrl D进行添加")
            }
        }
    })

    // const clientH = document.screen.height;
    const clientH = $(document).height();

    $('.body').height((clientH) + "px");

    i()

    function randomString(random_string_box_id, row, len) {
        row = row || 26
        len = len || 36

        let string_contain_number = $(random_string_box_id + ' input.string-contain-number').prop("checked")
        let string_contain_uppercase = $(random_string_box_id + ' input.string-contain-uppercase').prop("checked")
        let string_contain_lowercase = $(random_string_box_id + ' input.string-contain-lowercase').prop("checked")
        let string_contain_special = $(random_string_box_id + ' input.string-contain-special').prop("checked")
        let string_contain_take_out_confused = $(random_string_box_id + ' input.string-contain-take-out-confused').prop("checked")
        let number_chars, uppercase_chars, lowercase_chars, special_chars, chars = '';
        if (string_contain_take_out_confused) {
            number_chars = '0123456789';
            uppercase_chars = 'ABCDEFGHJKMNPQRSTUVWXYZ';
            lowercase_chars = 'abcdefghjkmnpqrstuvwxyz';
            special_chars = '~!@#$%^&*()_+';
        } else {
            number_chars = '0123456789';
            uppercase_chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
            lowercase_chars = 'abcdefghijklmnopqrstuvwxyz';
            lowercase_chars = 'abcdefghijklmnopqrstuvwxyz';
            special_chars = '~!@#$%^&*()_+';
        }

        // alert(string_contain_number)
        // console.log( $('#random-string-box-1 input.string-contain-number').prop("checked"))

        if (string_contain_number) {
            chars += number_chars
        }

        if (string_contain_uppercase) {
            chars += uppercase_chars
        }

        if (string_contain_lowercase) {
            chars += lowercase_chars
        }

        if (string_contain_special) {
            chars += special_chars
        }

        let maxPos = chars.length;
        let pwd = "";

        for (let i = 0; i < (len * row); i++) {
            pwd += chars.charAt(Math.floor(Math.random() * maxPos));
            if ((i + 1) % len == 0 && i != len * row) {
                pwd += "<br/>"
            }
        }

        return pwd;
    }

    function randomStringBox1() {
        $('#random-string-box-1 .string-box-p').html(randomString("#random-string-box-1", 10, 36));
    }

    function randomStringBox2() {
        $('#random-string-box-2 .string-box-p').html(randomString("#random-string-box-2", 10, 36));
    }

    function randomStringBox3() {
        $('#random-string-box-3 .string-box-p').html(randomString("#random-string-box-3", 10, 36));
    }

    function i() {
        randomStringBox1()
        randomStringBox2()
        randomStringBox3()
    }

</script>
{{template "layout_footer.tpl" .}}
