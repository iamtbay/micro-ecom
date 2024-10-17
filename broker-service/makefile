AUTH_SERVICE_DIR:= ../authentication-service
PRODUCTS_SERVICE_DIR:= ../products-service
CART_SERVICE_DIR:= ../cart-service
ORDERS_SERVICE_DIR:= ../orders-service
BROKER_SERVICE_DIR:= ../broker-service

build: 
	go build -o brokerService ./cmd/*.go

run: build
	./brokerService &

build_auth_service:
	@echo "Building auth service..."
	cd $(AUTH_SERVICE_DIR) && go build -o authService cmd/*.go

build_products_service:
	@echo "Building products service..."
	cd $(PRODUCTS_SERVICE_DIR) && go build -o productsService cmd/*.go

build_cart_service:
	@echo "Building cart service..."
	cd $(CART_SERVICE_DIR) && go build -o cartService cmd/*.go

build_orders_service:
	@echo "Building orders service..."
	cd $(ORDERS_SERVICE_DIR) && go build -o ordersAddressesService cmd/*.go

buildAll: build build_auth_service build_products_service build_cart_service build_orders_service
	@echo "All services built succesfully"

# run
run_auth_service:
	cd $(AUTH_SERVICE_DIR) && ./authService &

run_products_service:
	cd $(PRODUCTS_SERVICE_DIR) && ./productsService &

run_cart_service:
	cd $(CART_SERVICE_DIR) && ./cartService &

run_orders_service:
	cd $(ORDERS_SERVICE_DIR) && ./ordersAddressesService &

runAll: run_auth_service run_products_service run_cart_service run_orders_service run
	@echo "All services running"

stopAll:
	-@pkill -f authService || true
	-@pkill -f productsService || true
	-@pkill -f cartService || true
	-@pkill -f ordersAddressesService || true
	-@pkill -f brokerService || true
	@echo "All services stopped"