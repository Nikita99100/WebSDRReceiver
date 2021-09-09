require 'faye/websocket'
require 'eventmachine'

frequency = "1010"
band      = "0"
pcm_output = nil
EM.run {
  output = nil

  ws = Faye::WebSocket::Client.new('ws://websdr.ewi.utwente.nl:8901/~~stream?v=11', [], headers: {
    'Connection' => 'Upgrade',
    'Host' => 'websdr.ewi.utwente.nl:8901',
    'Origin' => 'http://websdr.ewi.utwente.nl:8901',
    'Cookie' => 'view=2; ID=5d83b2fd9c13',
    'Pragma' => 'no-cache',
    'Cache-Control' => 'no-cache',
    'Sec-WebSocket-Key' => 'K+3pZ6WnO+JMDVF5CX4J+w==',
    'Sec-WebSocket-Version' => '13',
    'Sec-WebSocket-Extensions' => 'x-webkit-deflate-frame',
    'User-Agent' => 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0 Safari/605.1.15'
  })

  ws.on :open do |event|
    puts "WebSocket connected, sending params to set the tuner..."
    pcm_output = File.open('pcm', 'wb')
    ws.send("GET /~~param?f=#{frequency}&band=#{band}&lo=-4dd.5&hi=4.5&mode=1&name=")
  end

  ws.on :message do |event|
    puts "Received #{event.data.size} bytes"
    event.data.each do |byte|
      pcm_output.print byte.chr
    end
    # p [:message, event.data]
  end

  ws.on :close do |event|
    pcm_output.close
    puts "Closed."
    ws = nil
  end
}