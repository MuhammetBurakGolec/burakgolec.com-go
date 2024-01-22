# Site
<a href="https://burakgolec.com/">burakgolec.com</a>

# Purpose
This is a simple website that I created to learn how to use HTML, CSS, JavaScript And Go. I also used this website to learn how to use Nginx.

# How to run
1. Install Go
2. Install Nginx
3. Clone this repository
4. Run `go run *.go` in the root directory
5. Set Configurations in Nginx `/etc/nginx/sites-available/default``
6. Set link to the Nginx `/etc/nginx/sites-enabled/default` 
    ```bash 
    sudo ln -s /etc/nginx/sites-available/default /etc/nginx/sites-enabled/default
    ```
7. Run `sudo nginx -s reload` in the root directory

# !!! Attention !!!
This website is not finished yet. I will add more features in the future.

This website is not secure. I will add SSL in the future.