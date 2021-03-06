# secrun
Run bash command as frequently as every second (a bit like cron)

## Use
  secrun --help
  (Displays help)

  secrun --version
  (Displays version)

  secrun --cmd mybash.sh
  (executes bash mybash.sh every second by default)
  
  secrun --cmd threebash.sh --freq 3
  (executes bash threebash.sh every three seconds)
  
  secrun --cmd hourmax.sh --freq 60 --max 3600
  (executes bash hourmax.sh approx every 60 seconds for an hour)

## Notes
  The timing of command is roughly a second apart but might be fractionally more. If hi-res timing is required, do not rely on freq
  
## To Do/ideas
  * More flexible command execution, multiple parameters etc.
  * Improve Timer resolution to synchronise with seconds
  * Create Blocking mode option (previous command must finish before executing again). For now can just be implemented with tmp file creation at start of bash script and its removal at end of script, with bash script termination if tmp file present.
