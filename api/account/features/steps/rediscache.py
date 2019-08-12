#!/usr/bin/env python3

from behave import *
from hamcrest import *
import json


@given('rediscache.token 创建 token: "{token:str}", phone: "{phone:str}", email: "{email:str}", password: "{password:str}", firstname: "{firstname:str}", lastname: "{lastname:str}", birthday: "{birthday:str}", gender: {gender:int}')
def step_impl(context, token, phone, email, password, firstname, lastname, birthday, gender):
    account = {
        "phone": phone,
        "email": email,
        "password": password,
        "firstName": firstname,
        "lastName": lastname,
        "birthday": birthday + "T00:00:00+08:00",
        "gender": gender,
    }
    context.redis_client.set(token, json.dumps(account))


@then('检查 rediscache.token, 存在记录 phone: "{phone:str}", email: "{email:str}", password: "{password:str}", firstname: "{firstname:str}", lastname: "{lastname:str}", birthday: "{birthday:str}", gender: {gender:int}')
def step_impl(context, phone, email, password, firstname, lastname, birthday, gender):
    res = context.redis_client.get(context.res["token"])
    account = json.loads(res)
    assert_that(account["phone"], equal_to(phone))
    assert_that(account["email"], equal_to(email))
    assert_that(account["password"], equal_to(password))
    assert_that(account["firstName"], equal_to(firstname))
    assert_that(account["lastName"], equal_to(lastname))
    assert_that(account["birthday"][:10], equal_to(birthday))
    assert_that(account["gender"], equal_to(gender))
