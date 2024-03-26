import re
import datetime  # eval()时含有datetime类型，不能去掉


def log2sql(log):
    """日志转sql"""
    datetime_pattern = "(\d{4}-\d{1,2}-\d{1,2}\s\d{1,2}:\d{1,2}:\d{1,2})"  # 时间字符串正则表达式
    datetime_repl = lambda x: '"{}"'.format(x.group())
    result = []
    for s in log.strip().split("\n"):
        sql, parameters = eval(s)
        for parameter in parameters:
            sql = sql.replace("?", str(parameter), 1)
        sql = sql.replace("None", "NULL")
        sql = re.sub(pattern=datetime_pattern, repl=datetime_repl, string=sql)  # 给时间字符串加上双引号
        result.append(sql)
    return result


if __name__ == '__main__':
    log = """
('SELECT "t1"."id", MAX("t1"."create_at") AS "create_at" FROM "product" AS "t1" GROUP BY "t1"."type"', [])
('SELECT "t1"."id", "t1"."name", "t1"."type", "t1"."create_at" FROM "product" AS "t1" INNER JOIN (SELECT "t1"."id", MAX("t1"."create_at") AS "create_at" FROM "product" AS "t1" GROUP BY "t1"."type") AS "lastest" ON (("t1"."id" = "lastest"."id") AND ("t1"."create_at" = "lastest"."create_at")) WHERE (((? AND ("t1"."create_at" IS NOT ?)) AND "t1"."create_at") != ?)', [datetime.datetime(2021, 1, 1, 0, 0), None, datetime.datetime(1, 1, 1, 0, 0)])
"""
    result = log2sql(log)
    for sql in result:
        print(sql)
    print('共{}条'.format(len(result)))
