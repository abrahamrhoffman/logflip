
Logflip v0.1

Nov 6th, 2015

Why Logflip? 

Logflip is a logrotate replacement. If you have used logrotate in the past,
then you know how effective it can be. However, logrotate starts to break  
down once you need to run it more than daily. This is because logrotate is 
old. The original creators didn't take into account that logs could grow at
a rate of hundreds of megabytes per second. Or maybe the designer just     
didn't have a need to rotate any faster. Regardless, with the advent of    
Cloud, Engineers and Operations teams have a critical need to rotate logs  
as quickly as possible under specific conditions.                          

LOGFLIP USAGE
Logflip is easy. It installs itself. Just run 'sudo ./logflip'
After that, you can 'service logflip start|stop|status' it.   
- Creates (if none exists) /var/logflip/logflip.conf          
- Automagically scans your OS for .log files and adds         
  to logflip.conf                                             
- Constantly monitors .log files and "flips" them when        
  they reach 100MB in size. file.log -> file.log.1            
- When .log reaches 100MB in size & .log.1 reaches 100MB      
  in size, logflip deletes .log.1.                            
   That's it! New features to come with new releases.        
