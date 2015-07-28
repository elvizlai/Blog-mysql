/**
 * Created by elvizlai on 15/3/23.
 */
$().ready(function () {
        //为input添加autocomplete
        $('input:not([autocomplete]),textarea:not([autocomplete]),select:not([autocomplete])').attr('autocomplete', 'off');

        //表单验证
        $('.validate').validate({
            rules: {
                email: {
                    required: true,
                    email: true
                },
                nickname: {
                    required: true,
                    minlength: 5
                },
                password: {
                    required: true,
                    minlength: 6
                },
                repeatpassword: {
                    required: true,
                    equalTo: '#password'
                }
            },
            messages: {
                email: {
                    required: "Email should not be null",
                    email: "Must be an email"
                },
                password: {
                    required: "Password should not be null",
                    minlength: "Password too short"
                }
            },
            errorPlacement: function (error, element) {
                error.appendTo(element.parent().parent());
            }
        });

        //登陆
        $('#loginSubmit').on('click', function () {
            if (!$('#loginForm').valid()) {
                return
            }

            var request = {
                email: $('#email').val(),
                password: $('#password').val()
            };

            var jsonStr = JSON.stringify(request);
            $.post('/login', jsonStr, function (data) {
                if (data.code == 0) {
                    window.location.href = "/articleList"
                } else {
                    var $loginError = $('#loginError');
                    $loginError.parent().show();
                    $loginError.text(data.msg);
                }
            });
        });

        //注册
        $('#registerSubmit').on('click', function () {
            if (!$('#registerForm').valid()) {
                return
            }
            var request = {
                email: $('#email').val(),
                nickname: $('#nickname').val(),
                password: $('#password').val()
            };

            var jsonStr = JSON.stringify(request);
            $.post('/register', jsonStr, function (data) {
                if (data.code == 0) {
                    window.location.href = "/login"
                } else {
                    var $registerError = $('#registerError');
                    $registerError.parent().show();
                    $registerError.text(data.msg);
                }
            });

        });

        //点击div button时联动a标签
        $('.button').on('click', function () {
            var href = $(this).children('a').attr('href');
            if (href != undefined) {
                window.location.href = href;
            }
        });


        $('.dropdown')
            .dropdown({
                transition: 'drop'
            });

        $('.combo')
            .dropdown({
                action: 'combo'
            }
        );

        //代码高亮
        $('pre').each(function () {
            $(this).addClass("line-numbers");
            var langClass = $(this).attr('class').split(";")[0].split(":")[1];
            $(this).wrapInner("<code class='language-" + langClass + "'></code>")
            var lines = $(this).text().split('\n').length;
            var $numbering = $('<ul/>').addClass('pre-numbering');
            $(this).append($numbering);
            for (i = 1; i <= lines; i++) {
                $numbering.append($('<li/>').text(i));
            }
        });

        $(".menu .item").click(function () {
            if (!$(this).hasClass("header")) {
                $(this)
                    .addClass('active')
                    .siblings()
                    .removeClass('active')
                ;
            }
        });

        $('.message .close').on('click', function () {
            $(this).closest('.message').fadeOut();
        });

        $(".Login").click(function () {
            window.open("/login", "_self");
        });
        $(".Register").click(function () {
            window.open("/register", "_self");
        });

        $(".ArticleList").click(function () {
            window.open("/articleList", "_self");
        });

        $(".Logout").click(function () {
            window.open("/logout", "_self");
        });

        //新建分类
        $('#categoryId').change(function () {
            if ($(this).children('option:selected').val() == -1) {
                newCategory();
            }
        });

        //修改分类
        $('#modifyCategory').click(function () {
            var id = $('#categoryId option:selected').val();
            if (id) {
                modifyCategory(id)
            } else {
                alert("请先选中要编辑的分类");
            }

        });
    }
);

//新建分类
function newCategory() {
    var categoryId;
    $('#newCategoryModel')
        .modal({
            closable: false,
            onDeny: function () {
                //将select置空
                $('#categoryId').parent().dropdown('restore defaults')
            },
            onApprove: function () {
                var nameObject = $('#newCategoryName');
                if (nameObject.val() == '') {
                    nameObject.parent().addClass('error');
                } else {
                    //ajax提交
                    $.post("/addCategory", {
                            categoryName: nameObject.val(),
                            _xsrf: $('meta[name=_xsrf]').attr('content')
                        },
                        function (data) {
                            if (data.errcode) {
                                //返回值不为0-->失败
                                nameObject.parent().addClass('error');
                            } else {
                                //成功
                                nameObject.val('');
                                categoryId = data.result.Id;
                                $('#categoryId').prepend('<option value="' + data.result.Id + '">' + data.result.Name + '</option>');
                                $('#newCategoryModel').modal('hide');
                            }
                        }
                    );
                }
                return false;
            },
            onHidden: function () {
                if (categoryId) {
                    $('#categoryId').parent().dropdown('set selected', categoryId);
                }
            }
        }).modal('show');
}

//修改分类
function modifyCategory(id) {


    $('#modifyCategoryName').val($('#categoryId').parent().find('.item.active').text());

    $('#modifyCategoryModel')
        .modal({
            closable: false,
            onDeny: function () {

            },
            onApprove: function () {
                var nameObject = $('#modifyCategoryName');
                if (nameObject.val() == "") {
                    nameObject.parent().addClass('error');
                } else {
                    //ajax提交
                    $.post("/modifyCategory", {
                            categoryId: id,
                            categoryName: nameObject.val(),
                            _xsrf: $('meta[name=_xsrf]').attr('content')
                        },
                        function (data) {
                            if (data.errcode) {
                                //返回值不为0-->失败
                                nameObject.parent().addClass('error');
                            } else {
                                var selObj = $('#categoryId').parent();
                                selObj.dropdown('set text', data.result.Name);
                                selObj.find('option[value="' + id + '"]').text(data.result.Name);
                                selObj.find('.active').text(data.result.Name);
                                $('#modifyCategoryModel').modal('hide');
                            }
                        }
                    );
                }
                return false;
            }
        }).modal('show');
}

//文章提交
function postArticle() {

}


//表单验证
function formValid() {
    $('.verifyForm').form({
        email: {
            identifier: 'email',
            rules: [
                {
                    type: 'email',
                    prompt: 'Please enter an email'
                }
            ]
        },
        password: {
            identifier: 'password',
            rules: [
                {
                    type: 'length[6]',
                    prompt: 'Your password must be at least 6 characters'
                }
            ]
        },
        confirmpassword: {
            identifier: 'confirmpassword',
            rules: [
                {
                    type: 'empty',
                    prompt: 'Please enter a confirmPassword'
                },
                {
                    type: 'match[password]',
                    prompt: 'Your password must be same as you input before'
                }
            ]
        },
        articleTitle: {
            identifier: 'articleTitle',
            rules: [
                {
                    type: 'empty',
                    prompt: 'Please enter an articleTitle'
                }
            ]
        },
        categoryName: {
            identifier: 'categoryName',
            rules: [
                {
                    type: 'empty',
                    prompt: 'Please enter an categoryName'
                }
            ]
        },
        categoryId: {
            identifier: 'categoryId',
            rules: [
                {
                    type: 'empty',
                    prompt: 'Please enter an categoryId'
                }
            ]
        }
    });
}
