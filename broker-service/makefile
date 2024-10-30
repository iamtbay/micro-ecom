AUTH_SERVICE_DIR:= ../authentication-service
PRODUCTS_SERVICE_DIR:= ../products-service
CART_SERVICE_DIR:= ../cart-service
ORDERS_SERVICE_DIR:= ../orders-service
REVIEWS_SERVICE_DIR:= ../reviews-service
BROKER_SERVICE_DIR:= ../broker-service
INVENTORY_SERVICE_DIR:= ../inventory-service

build: 
	go build -o brokerService ./cmd/*.go

run: build
	./brokerService &

build_inventory_service:
	@echo "Building inventory service..."
	cd $(INVENTORY_SERVICE_DIR) && go build -o inventoryService cmd/*.go
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

build_reviews_service:
	@echo "Building reviews service..."
	cd $(REVIEWS_SERVICE_DIR) && go build -o reviewsService cmd/*.go

buildAll: build build_auth_service build_products_service build_cart_service build_orders_service build_reviews_service build_inventory_service
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

run_reviews_service:
	cd $(REVIEWS_SERVICE_DIR) && ./reviewsService &

run_inventory_service:
	cd $(INVENTORY_SERVICE_DIR) && ./inventoryService &

stopAll:
	-@pkill -f authService || true
	-@pkill -f productsService || true
	-@pkill -f cartService || true
	-@pkill -f ordersAddressesService || true
	-@pkill -f brokerService || true
	-@pkill -f reviewsService || true
	-@pkill -f inventoryService || true
	@echo "All services stopped"

runAll: stopAll run_auth_service run_products_service run_cart_service run_orders_service run_reviews_service run_inventory_service run
	@echo "All services running"
