# Reference https://en.wikipedia.org/wiki/Trie
# Example
# t = Trie.new
# t.insert("hello")
# t.search("hell") => false
# t.starts_with("hell") => true
class Trie
  attr_accessor :children, :is_word

  def initialize()
    @children = {}
    @is_word = false
  end

  def insert(word)
    children = @children
    word.each_char.each_with_index do |char, index|
      children[char] ||= Trie.new
      children[char].is_word = true if word.size == index + 1
      children = children[char].children
    end
  end

  # Returns true if the word is in the trie.
  def search(word)
    match_helper(word, true)
  end

  # Returns true if there is any word in the trie that starts with the given prefix.
  def starts_with(prefix)
    match_helper(prefix)
  end

  private

  def match_helper(match, is_word = false)
    children = @children
    found = false
    match.each_char.each_with_index do |char, index|
      break if children[char].nil?
      if match.size == index + 1 && requires_word?(is_word, children[char].is_word)
        found = true
      end
      children = children[char].children
    end
    found
  end

  def requires_word?(requires_word, node_is_word)
    return true unless requires_word
    node_is_word
  end
end
