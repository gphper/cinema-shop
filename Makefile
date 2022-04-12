# cinema-shop makefile

#         service-method-type 
# eg make usercenter-run-api
#    make usercenter-gen-api

ROOTPATH=./services/
SERVERNAME = $(word 1,$(subst -, ,${MAKECMDGOALS}))
METHOD = $(word 2,$(subst -, ,${MAKECMDGOALS}))
TYPE = $(word 3,$(subst -, ,${MAKECMDGOALS}))
SERVERPATH = $(ROOTPATH)$(SERVERNAME)
%:
ifeq ($(METHOD),run)
ifeq ($(TYPE),api)
	go run $(SERVERPATH)/api/$(SERVERNAME).go -f $(SERVERPATH)/api/etc/$(SERVERNAME)-api.yaml
else
	go run $(SERVERPATH)/rpc/$(SERVERNAME).go -f $(SERVERPATH)/rpc/etc/$(SERVERNAME).yaml
endif
else
ifeq ($(TYPE),api)
	goctl api go --api $(SERVERPATH)/api/desc/$(SERVERNAME).api --dir $(SERVERPATH)/api --style goZero
else
	goctl rpc protoc $(SERVERPATH)/rpc/pb/$(SERVERNAME)/$(SERVERNAME).proto --go_out=$(SERVERPATH)/rpc/pb \
 	 --go-grpc_out=$(SERVERPATH)/rpc/pb --zrpc_out=$(SERVERPATH)/rpc --style=goZero
endif
endif