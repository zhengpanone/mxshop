# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: goods.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0bgoods.proto\x1a\x1bgoogle/protobuf/empty.proto\"\xbe\x01\n\x12GoodsFilterRequest\x12\x11\n\tprice_min\x18\x01 \x01(\x05\x12\x11\n\tprice_max\x18\x02 \x01(\x05\x12\x0e\n\x06is_hot\x18\x03 \x01(\x08\x12\x0e\n\x06is_new\x18\x04 \x01(\x08\x12\x0e\n\x06is_tab\x18\x05 \x01(\x08\x12\x14\n\x0ctop_category\x18\x06 \x01(\x05\x12\x0c\n\x04page\x18\x07 \x01(\r\x12\x0c\n\x04size\x18\x08 \x01(\r\x12\x11\n\tkey_words\x18\t \x01(\t\x12\r\n\x05\x62rand\x18\n \x01(\x05\"D\n\x11GoodsListResponse\x12\r\n\x05total\x18\x01 \x01(\x05\x12 \n\x04\x64\x61ta\x18\x02 \x03(\x0b\x32\x12.GoodsInfoResponse\"\x1e\n\x10\x42\x61tchGoodsIdInfo\x12\n\n\x02id\x18\x01 \x03(\x05\"\xb9\x02\n\x0f\x43reateGoodsInfo\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x10\n\x08goods_sn\x18\x03 \x01(\t\x12\x0e\n\x06stocks\x18\x07 \x01(\x05\x12\x14\n\x0cmarket_price\x18\x08 \x01(\x02\x12\x12\n\nshop_price\x18\t \x01(\x02\x12\x13\n\x0bgoods_brief\x18\n \x01(\t\x12\x12\n\ngoods_desc\x18\x0b \x01(\t\x12\x11\n\tship_free\x18\x0c \x01(\x08\x12\x0e\n\x06images\x18\r \x03(\t\x12\x13\n\x0b\x64\x65sc_images\x18\x0e \x03(\t\x12\x19\n\x11goods_front_image\x18\x0f \x01(\t\x12\x0e\n\x06is_new\x18\x10 \x01(\x08\x12\x0e\n\x06is_hot\x18\x11 \x01(\x08\x12\x0f\n\x07on_sale\x18\x12 \x01(\x08\x12\x13\n\x0b\x63\x61tegory_id\x18\x13 \x01(\x05\"5\n\x19\x43\x61tegoryBriefInfoResponse\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0c\n\x04name\x18\x02 \x01(\t\"\x13\n\x11\x42randInfoResponse\"\xc4\x03\n\x11GoodsInfoResponse\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x13\n\x0b\x63\x61tegory_id\x18\x02 \x01(\x05\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x10\n\x08goods_sn\x18\x04 \x01(\t\x12\x11\n\tclick_num\x18\x05 \x01(\x05\x12\x10\n\x08sold_num\x18\x06 \x01(\x05\x12\x0f\n\x07\x66\x61v_num\x18\x07 \x01(\x05\x12\x14\n\x0cmarket_price\x18\t \x01(\x02\x12\x12\n\nshop_price\x18\n \x01(\x02\x12\x13\n\x0bgoods_brief\x18\x0b \x01(\t\x12\x12\n\ngoods_desc\x18\x0c \x01(\t\x12\x11\n\tship_free\x18\r \x01(\x08\x12\x0e\n\x06images\x18\x0e \x03(\t\x12\x13\n\x0b\x64\x65sc_images\x18\x0f \x03(\t\x12\x19\n\x11goods_front_image\x18\x10 \x01(\t\x12\x0e\n\x06is_new\x18\x11 \x01(\x08\x12\x0e\n\x06is_hot\x18\x12 \x01(\x08\x12\x0f\n\x07on_sale\x18\x13 \x01(\x08\x12\x10\n\x08\x61\x64\x64_time\x18\x14 \x01(\x03\x12,\n\x08\x63\x61tegory\x18\x15 \x01(\x0b\x32\x1a.CategoryBriefInfoResponse\x12!\n\x05\x62rand\x18\x16 \x01(\x0b\x32\x12.BrandInfoResponse\"\x1d\n\x0f\x44\x65leteGoodsInfo\x12\n\n\x02id\x18\x01 \x01(\x05\"\x1d\n\x0fGoodInfoRequest\x12\n\n\x02id\x18\x01 \x01(\x05\"9\n\x1b\x43\x61tegoryFilterFilterRequest\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0e\n\x06is_tab\x18\x02 \x01(\x08\"]\n\x14\x43\x61tegoryListResponse\x12\r\n\x05total\x18\x01 \x01(\x05\x12#\n\x04\x64\x61ta\x18\x02 \x03(\x0b\x32\x15.CategoryInfoResponse\x12\x11\n\tjson_data\x18\x03 \x01(\t\"\x15\n\x13\x43\x61tegoryListRequest\"\x7f\n\x17SubCategoryListResponse\x12\r\n\x05total\x18\x01 \x01(\x05\x12#\n\x04info\x18\x02 \x01(\x0b\x32\x15.CategoryInfoResponse\x12\x30\n\x11sub_category_list\x18\x03 \x03(\x0b\x32\x15.CategoryInfoResponse\"h\n\x14\x43\x61tegoryInfoResponse\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x17\n\x0fparent_category\x18\x03 \x01(\x05\x12\r\n\x05level\x18\x04 \x01(\x05\x12\x0e\n\x06is_tab\x18\x05 \x01(\x08\"\x14\n\x12\x44\x65leteCategoryInfo\"0\n\x14QueryCategoryRequest\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0c\n\x04name\x18\x02 \x01(\t\"N\n\x18\x42\x61tchCategoryInfoRequest\x12\n\n\x02id\x18\x01 \x03(\x05\x12\x12\n\ngoods_nums\x18\x02 \x01(\x05\x12\x12\n\nbrand_nums\x18\x03 \x01(\x05\"\x13\n\x11\x42randListResponse\"\x0e\n\x0c\x42randRequest\"\x14\n\x12\x42\x61nnerListResponse\"\x0f\n\rBannerRequest\"\x10\n\x0e\x42\x61nnerResponse\"\x1c\n\x1a\x43\x61tegoryBrandFilterRequest\"\x1b\n\x19\x43\x61tegoryBrandListResponse\"g\n\x13\x43\x61tegoryInfoRequest\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x17\n\x0fparent_category\x18\x03 \x01(\x05\x12\r\n\x05level\x18\x04 \x01(\x05\x12\x0e\n\x06is_tab\x18\x05 \x01(\x08\"\x16\n\x14\x43\x61tegoryBrandRequest\"\x17\n\x15\x43\x61tegoryBrandResponse2\xae\x0b\n\x05Goods\x12\x34\n\tGoodsList\x12\x13.GoodsFilterRequest\x1a\x12.GoodsListResponse\x12\x36\n\rBatchGetGoods\x12\x11.BatchGoodsIdInfo\x1a\x12.GoodsListResponse\x12\x33\n\x0b\x43reateGoods\x12\x10.CreateGoodsInfo\x1a\x12.GoodsInfoResponse\x12\x37\n\x0b\x44\x65leteGoods\x12\x10.DeleteGoodsInfo\x1a\x16.google.protobuf.Empty\x12\x37\n\x0bUpdateGoods\x12\x10.CreateGoodsInfo\x1a\x16.google.protobuf.Empty\x12\x36\n\x0eGetGoodsDetail\x12\x10.GoodInfoRequest\x1a\x12.GoodsInfoResponse\x12\x43\n\x12GetAllCategoryList\x12\x16.google.protobuf.Empty\x1a\x15.CategoryListResponse\x12@\n\x0eGetSubCategory\x12\x14.CategoryListRequest\x1a\x18.SubCategoryListResponse\x12=\n\x0e\x43reateCategory\x12\x14.CategoryInfoRequest\x1a\x15.CategoryInfoResponse\x12=\n\x0e\x44\x65leteCategory\x12\x13.DeleteCategoryInfo\x1a\x16.google.protobuf.Empty\x12>\n\x0eUpdateCategory\x12\x14.CategoryInfoRequest\x1a\x16.google.protobuf.Empty\x12\x37\n\tBrandList\x12\x16.google.protobuf.Empty\x1a\x12.BrandListResponse\x12\x30\n\x0b\x43reateBrand\x12\r.BrandRequest\x1a\x12.BrandInfoResponse\x12\x34\n\x0b\x44\x65leteBrand\x12\r.BrandRequest\x1a\x16.google.protobuf.Empty\x12\x34\n\x0bUpdateBrand\x12\r.BrandRequest\x1a\x16.google.protobuf.Empty\x12\x39\n\nBannerList\x12\x16.google.protobuf.Empty\x1a\x13.BannerListResponse\x12/\n\x0c\x43reateBanner\x12\x0e.BannerRequest\x1a\x0f.BannerResponse\x12\x36\n\x0c\x44\x65leteBanner\x12\x0e.BannerRequest\x1a\x16.google.protobuf.Empty\x12\x36\n\x0cUpdateBanner\x12\x0e.BannerRequest\x1a\x16.google.protobuf.Empty\x12L\n\x11\x43\x61tegoryBrandList\x12\x1b.CategoryBrandFilterRequest\x1a\x1a.CategoryBrandListResponse\x12@\n\x14GetCategoryBrandList\x12\x14.CategoryInfoRequest\x1a\x12.BrandListResponse\x12\x44\n\x13\x43reateCategoryBrand\x12\x15.CategoryBrandRequest\x1a\x16.CategoryBrandResponse\x12\x44\n\x13\x44\x65leteCategoryBrand\x12\x15.CategoryBrandRequest\x1a\x16.google.protobuf.Empty\x12\x44\n\x13UpdateCategoryBrand\x12\x15.CategoryBrandRequest\x1a\x16.google.protobuf.EmptyB\tZ\x07.;protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'goods_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\007.;proto'
  _globals['_GOODSFILTERREQUEST']._serialized_start=45
  _globals['_GOODSFILTERREQUEST']._serialized_end=235
  _globals['_GOODSLISTRESPONSE']._serialized_start=237
  _globals['_GOODSLISTRESPONSE']._serialized_end=305
  _globals['_BATCHGOODSIDINFO']._serialized_start=307
  _globals['_BATCHGOODSIDINFO']._serialized_end=337
  _globals['_CREATEGOODSINFO']._serialized_start=340
  _globals['_CREATEGOODSINFO']._serialized_end=653
  _globals['_CATEGORYBRIEFINFORESPONSE']._serialized_start=655
  _globals['_CATEGORYBRIEFINFORESPONSE']._serialized_end=708
  _globals['_BRANDINFORESPONSE']._serialized_start=710
  _globals['_BRANDINFORESPONSE']._serialized_end=729
  _globals['_GOODSINFORESPONSE']._serialized_start=732
  _globals['_GOODSINFORESPONSE']._serialized_end=1184
  _globals['_DELETEGOODSINFO']._serialized_start=1186
  _globals['_DELETEGOODSINFO']._serialized_end=1215
  _globals['_GOODINFOREQUEST']._serialized_start=1217
  _globals['_GOODINFOREQUEST']._serialized_end=1246
  _globals['_CATEGORYFILTERFILTERREQUEST']._serialized_start=1248
  _globals['_CATEGORYFILTERFILTERREQUEST']._serialized_end=1305
  _globals['_CATEGORYLISTRESPONSE']._serialized_start=1307
  _globals['_CATEGORYLISTRESPONSE']._serialized_end=1400
  _globals['_CATEGORYLISTREQUEST']._serialized_start=1402
  _globals['_CATEGORYLISTREQUEST']._serialized_end=1423
  _globals['_SUBCATEGORYLISTRESPONSE']._serialized_start=1425
  _globals['_SUBCATEGORYLISTRESPONSE']._serialized_end=1552
  _globals['_CATEGORYINFORESPONSE']._serialized_start=1554
  _globals['_CATEGORYINFORESPONSE']._serialized_end=1658
  _globals['_DELETECATEGORYINFO']._serialized_start=1660
  _globals['_DELETECATEGORYINFO']._serialized_end=1680
  _globals['_QUERYCATEGORYREQUEST']._serialized_start=1682
  _globals['_QUERYCATEGORYREQUEST']._serialized_end=1730
  _globals['_BATCHCATEGORYINFOREQUEST']._serialized_start=1732
  _globals['_BATCHCATEGORYINFOREQUEST']._serialized_end=1810
  _globals['_BRANDLISTRESPONSE']._serialized_start=1812
  _globals['_BRANDLISTRESPONSE']._serialized_end=1831
  _globals['_BRANDREQUEST']._serialized_start=1833
  _globals['_BRANDREQUEST']._serialized_end=1847
  _globals['_BANNERLISTRESPONSE']._serialized_start=1849
  _globals['_BANNERLISTRESPONSE']._serialized_end=1869
  _globals['_BANNERREQUEST']._serialized_start=1871
  _globals['_BANNERREQUEST']._serialized_end=1886
  _globals['_BANNERRESPONSE']._serialized_start=1888
  _globals['_BANNERRESPONSE']._serialized_end=1904
  _globals['_CATEGORYBRANDFILTERREQUEST']._serialized_start=1906
  _globals['_CATEGORYBRANDFILTERREQUEST']._serialized_end=1934
  _globals['_CATEGORYBRANDLISTRESPONSE']._serialized_start=1936
  _globals['_CATEGORYBRANDLISTRESPONSE']._serialized_end=1963
  _globals['_CATEGORYINFOREQUEST']._serialized_start=1965
  _globals['_CATEGORYINFOREQUEST']._serialized_end=2068
  _globals['_CATEGORYBRANDREQUEST']._serialized_start=2070
  _globals['_CATEGORYBRANDREQUEST']._serialized_end=2092
  _globals['_CATEGORYBRANDRESPONSE']._serialized_start=2094
  _globals['_CATEGORYBRANDRESPONSE']._serialized_end=2117
  _globals['_GOODS']._serialized_start=2120
  _globals['_GOODS']._serialized_end=3574
# @@protoc_insertion_point(module_scope)