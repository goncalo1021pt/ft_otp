NAME = ft_otp

SRCS = $(wildcard $(SRCS_DIR)*.go)
SRCS_DIR = scrs/

GOFLAGS = 

#color codes
GREEN = \033[0;32m
RED = \033[0;31m
BLUE = \033[0;34m
ORANGE = \033[0;33m
NC = \033[0m

all: $(NAME)

$(NAME): $(SRCS)
	@echo "$(GREEN)$(NAME)$(NC) Building..."
	@go build $(GOFLAGS) -o $(NAME) $(SRCS)
	@echo "$(GREEN)$(NAME)$(NC) Build complete!"

clean:
	@go clean

fclean: clean
	@rm -f $(NAME)
	@rm -f ft_otp.key
	@rm -f qrcode.png
	@echo "$(RED)$(NAME)$(NC) Removed."

re: fclean all

test-g:
	./$(NAME) -g key.hex

test-k:
	./$(NAME) -k ft_otp.key

.PHONY: all clean fclean re test-g test-k