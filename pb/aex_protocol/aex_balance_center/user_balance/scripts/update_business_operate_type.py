#! /usr/bin/python3

import json

# 需要被修改的 proto 文件
generalProtoFile = "../proto/general.proto"

# 从此文件中读取最新的数据
generalProtoSrcFile = "../src/business_operate_type.json"

# 读取需要被修改的 proto 文件
general_proto = open(generalProtoFile)
general_proto_lines = general_proto.readlines()
general_proto.close()

# 读取 source file
source_file = open(generalProtoSrcFile)
source_json = json.load(source_file)
source_file.close()


lines_without_operate_types = []
lines_aex_business_types = []

into_aex_business = False
can_collect_business_type = False
aex_business_type_line = -1

class BusinessType():
    def __init__(self, business_id, enum, note='', is_valid=True):
        self.id = int(business_id)
        self.enum = enum
        self.comments = []
        self.note = note
        self.is_valid = is_valid

    def __str__(self):
        return "id: {}, emun: {}, comments: {}, note: {}".format(self.id, self.enum, self.comments, self.note)

    def add_comment(self, comment):
        self.comments.append(comment)

    def gen_lines(self):
        lines = []
        for comment in self.comments:
            lines.append("        // {}\n".format(comment))

        if self.note != '':
            lines.append("        // NOTE: {}\n".format(self.note))

        lines.append("        {} = {};\n".format(self.enum, self.id))
        return lines



def parse_business_type_line(line):
    elements = line.split("=")
    enum = elements[0].strip()
    business_id = elements[1].rstrip().rstrip(";").strip()
    business = BusinessType(business_id, enum)
    return business

def parse_business_type_lines(lines):
    comments = []
    businesses = []
    for line in lines:
        if is_note(line):
            comments.note = line.lstrip().lstrip("// NOTE:").strip()

        if is_comment(line):
            comments.append(line.lstrip().lstrip("//").strip())

        else:
            business = parse_business_type_line(line)
            for comment in comments:
                business.add_comment(comment)

            businesses.append(business)
            comments = []

    return businesses

def is_comment(line):
    return line.lstrip().startswith("//")

def is_note(line):
    return line.lstrip().startswith("// NOTE:")

for i, line in enumerate(general_proto_lines):
    # 如果已经进入到了业务类型这一行
    # 将这些行收集起来
    this_business = []
    if can_collect_business_type:
        if aex_business_type_line == -1:
            aex_business_type_line = i
        # 收集完毕
        if line !="    }\n":
            this_business.append(line)
            if is_comment(line) == False:
                this_business = []
            # 收集业务类型
            lines_aex_business_types.append(line)
            continue
        else:
            can_collect_business_type = False

    lines_without_operate_types.append(line)

    if into_aex_business:
        if can_collect_business_type == False:
            # 到这里，说明可以收集业务类型了
            if line == "    enum Type {\n":
                can_collect_business_type = True
                continue
    else:
        if line == 'message AexBusiness {\n':
            into_aex_business = True
        continue

# general_businesses = parse_business_type_lines(lines_aex_business_types)
general_businesses = []

for e in source_json:
    business = BusinessType(e.get('op_type_id'), e.get('define'), note=e.get('note'), is_valid=e.get('status')==1)
    business.comments.append(e.get('op_type'))
    general_businesses.append(business)
    
general_businesses.sort(key=lambda business: business.id)

final_file = lines_without_operate_types[:aex_business_type_line]
for business in general_businesses:
    final_file.extend(business.gen_lines())

final_file.extend(lines_without_operate_types[aex_business_type_line:])

general_new = open(generalProtoFile, 'w')
for line in final_file:
    general_new.write(line)
general_new.close()

# print("total lines:", len(general_proto_lines))
# print("lines without business type:", len(lines_without_operate_types))
# print("lines of business type:", len(lines_aex_business_types))

# for line in lines_without_operate_types:
#     print(line, end='')
