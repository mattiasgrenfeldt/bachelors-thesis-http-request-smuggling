# https://longliveruby.com/articles/rails-request-cycle
# https://github.com/rack/rack/blob/master/SPEC.rdoc
class ImitationOfRails
    def call(env)
        # puts(env)
        # env.each { |key, value| puts "k: #{key}, v: #{value}" }
        buf = env["rack.input"].read
        [200, {'Content-Type' => 'text/plain'}, ['HTTP_TRANSFER_ENCODING: ' + env.fetch("HTTP_TRANSFER_ENCODING", "") + ' CONTENT_LENGTH: ' + env.fetch("CONTENT_LENGTH", "") + ' Body length: ' + buf.length.to_s + ' Body: ' + buf]]
    end
end
