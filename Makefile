# cinema-shop makefile

#         service-method-type 
# eg make usercenter-run-api
#    make usercenter-gen-api

ROOTPATH= ./services/
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
endif

ifeq ($(METHOD),gen)
ifeq ($(TYPE),api)
	goctl api go --api $(SERVERPATH)/api/desc/$(SERVERNAME).api --home ./template  --dir $(SERVERPATH)/api --style goZero
else
	goctl rpc protoc $(SERVERPATH)/rpc/pb/$(SERVERNAME)/$(SERVERNAME).proto --home ./template --go_out=$(SERVERPATH)/rpc/pb \
 	 --go-grpc_out=$(SERVERPATH)/rpc/pb --zrpc_out=$(SERVERPATH)/rpc --style=goZero
endif
endif

ifeq ($(METHOD),new)
ifeq ($(TYPE),api)
	$(shell mkdir "$(SERVERPATH)/api/desc")
	goctl api -o $(SERVERPATH)/api/desc/$(SERVERNAME).api
else
	$(shell mkdir "$(SERVERPATH)/rpc/pb/$(SERVERNAME)")
	goctl rpc template -o $(SERVERPATH)/rpc/pb/$(SERVERNAME)/$(SERVERNAME).proto
endif
endif