import grpc

from userop_srv.proto import address_pb2,address_pb2_grpc


class AddressTest:
    pass
    # def __init__(self):
    #     channel = grpc.insecure_channel("127.0.0.1:5001")
    #     self.stub = user_pb2_grpc.UserStub(channel)
    #
    # def user_list(self):
    #     rsp:user_pb2.UserListResponse = self.stub.GetUserList(user_pb2.PageInfo())
    #     print(rsp.total)
    #     for user in rsp.data:
    #         print(user.mobile)
    #
    # def user_by_id(self,id):
    #     rsp:user_pb2.UserInfoResponse = self.stub.GetUserById(user_pb2.IdRequest(id=id))
    #     print(rsp)
    #
    # def user_by_mobile(self, mobile):
    #     rsp: user_pb2.UserInfoResponse = self.stub.GetUserByMobile(user_pb2.MobileRequest(mobile=mobile))
    #     print(rsp)



if __name__ == '__main__':
    pass
    # user = UserTest()
    # user.user_list()
    # user.user_by_id(10)
    # user.user_by_mobile('14548789993')
