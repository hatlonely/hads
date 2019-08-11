Feature: signup 注册测试

    Scenario Outline: 注册成功
        Given mysqldb.accounts 删除用户, email: "<email>"
        When 请求 /signup, phone: "<phone>", email: "<email>", password: "<password>", firstname: "<firstname>", secondname: "<secondname>", birthday: "<birthday>", gender: <gender>
        Then 检查状态码 res.status_code: <status>
        Then 检查注册返回包体 res.body, success: <success>
        Then 检查 mysqldb.accounts, 存在记录 phone: "<phone>", email: "<email>", password: "<password>", firstname: "<firstname>", secondname: "<secondname>", birthday: "<birthday>", gender: <gender>
        Examples:
            | phone       | email                  | password | firstname | secondname | birthday   | gender | status | success |
            | 13145678901 | hatlonely1@foxmail.com | 123456   | 贺        | 乐         | 1992-01-01 | 1      | 200    | true    |

    Scenario Outline: 重复注册
        Given mysqldb.accounts 删除用户, email: "<email>"
        Given mysqldb.accounts 创建用户, phone: "<phone>", email: "<email>", password: "<password>", firstname: "<firstname>", secondname: "<secondname>", birthday: "<birthday>", gender: <gender>
        When 请求 /signup, phone: "<phone>", email: "<email>", password: "<password>", firstname: "<firstname>", secondname: "<secondname>", birthday: "<birthday>", gender: <gender>
        Then 检查状态码 res.status_code: 500
        Then 检查返回包体 res.body，包含字符串 "phone [13145678901] is already exists"
        Examples:
            | phone       | email                  | password | firstname | secondname | birthday   | gender | status | success |
            | 13145678901 | hatlonely1@foxmail.com | 123456   | 贺        | 乐         | 1992-01-01 | 1      | 200    | true    |


    Scenario Outline: 异常注册
        When 请求 /signup, phone: "<phone>", email: "<email>", password: "<password>", firstname: "<firstname>", secondname: "<secondname>", birthday: "<birthday>", gender: <gender>
        Then 检查状态码 res.status_code: <status>
        Then 检查返回包体 res.body，包含字符串 "<body>"
        Examples:
            | phone        | email                                                             | password | firstname | secondname | birthday   | gender | status | body           |
            | 131-45678901 | hatlonely1@foxmail.com                                            | 123456   | 孙        | 悟空       | 1992-01-01 | 1      | 400    | 无效的电话号码 |
            | 13145678901  | hatlonely1                                                        | 123456   | 孙        | 悟空       | 1992-01-01 | 1      | 400    | 无效的邮箱     |
            | N/A          | N/A                                                               | 123456   | 孙        | 悟空       | 1992-01-01 | 1      | 400    | 必要字段       |
            | 13145678901  | veryveryveryveryveryveryveryveryveryveryveryvlongname@foxmail.com | 123456   | 孙        | 悟空       | 1992-01-01 | 1      | 400    | 至多64个字符   |
            | N/A          | hatlonely2@foxmail.com                                            | 123456   | 孙        | 悟空       | 1992-01-01 | 1      | 400    | 必要字段       |
            | 13145678903  | N/A                                                               | 123456   | 孙        | 悟空       | 1992-01-01 | 1      | 400    | 必要字段       |
