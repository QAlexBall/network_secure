export GODEBUG=x509ignoreCN=0

tmux new-session -d -s single 'go run server.go';
tmux split-window;
tmux send 'go run client.go' ENTER;
tmux a;