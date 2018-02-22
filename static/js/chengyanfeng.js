
//初始化插件
 Dropzone.autoDiscover = false;
    var myDropzone = new Dropzone("#myDropzone", {
        url: "/uplaod_user",
        addRemoveLinks: false,
        method: 'post',
        maxFiles:1,
        paramName:'file',
        filesizeBase: 1024,
        uploadMultiple:false,
        previewsContainer:null,
        maxFilesize:1,
        clickable:true,
        acceptedFiles:'.html',
        dictInvalidFileType: "你不能上传该类型文件,文件类型只能是*.jpg,*.gif,*.png。",
        dictFileTooBig:"文件过大上传文件最大支持.",
        sending: function(file, xhr, formData) {
         console.log( file);
        formData.append("size", file.size);
        formData.append("name", file.name);
        formData.append("total", "1");
        console.log(file.name)

},


       success: function (file, response, e) {
                if (code=200){
                setTimeout(function(){
                      //所有的逻辑
                      alert("html成功上传")
                },2000);
                }

        },
        complete:function(file){

        files=this.getRejectedFiles()
        for (var i=0;i<files.length;i++)
            {
           this.removeFile(files[i]);
            $("#tip").html("上传文件太大重新上传").css("color","red")
           }
        },

        maxfilesreached: function(files) {
        files=this.getRejectedFiles()
        for (var i=0;i<files.length;i++)
            {
           this.removeFile(files[i]);
            }
        },
        maxfilesexceeded:function(file){
        $("#tip").html("只能上传一个文件")
                this.removeFile(file);
        }


    });


$("#send").click(function(){

     var sub=$("#sub").val()
   if(sub==""){
    alert("请输入主题");
    return false
    }
        $.ajax({
        url:"/sendmail",
         dataType:"json",
          type:"post",
         data:{sub:sub},
        success:function(result){

        console.log(result.code)
            if (result.code==411){
            alert("发送失败，请上传用户文件")
            return false
            };
            if (result.code==412){
                        alert("发送失败，您还没有上传html")
                        return false
                        };
            if (result.code==413){
             alert("发送失败，您还没有上传html文件")
             return false
             };
             if (result.code==200){

                          alert("发送邮件成功")
                          location.reload()
                          };
        }});
    });

















$(function () {
           $("#btn_uploadimg").click(function () {

               var fileObj = document.getElementById("FileUpload").files[0]; // js 获取文件对象
               if (typeof (fileObj) == "undefined" || fileObj.size <= 0) {
                   alert("请选择图片");
                   return;
               }
               var formFile = new FormData();
               formFile.append("action", "UploadVMKImagePath");
               formFile.append("file", fileObj); //加入文件对象


             //第二种 ajax 提交

               var data = formFile;
               $.ajax({
                   url: "/uplaod_user",
                   data: data,
                   type: "Post",
                   dataType: "json",
                   cache: false,//上传文件无需缓存
                   processData: false,//用于对data参数进行序列化处理 这里必须false
                   contentType: false, //必须
                   success: function (result) {
                   $("#count").attr("value",result.count)
                       alert("上传完成!");
                   },
               })
           })
       })
