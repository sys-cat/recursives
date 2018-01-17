def recursive(n)
  if n==0
    return 1
  end
  return n * recursive(n-1)
end

print "#{recursive(100)}\n"
