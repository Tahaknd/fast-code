# FastCode: Snippet Yönetim Uygulaması

FastCode, geliştiriciler için hızlı ve verimli bir şekilde kod snippet'lerini (kod parçacıklarını) oluşturma, düzenleme, arama ve yönetme amacıyla geliştirilmiş bir masaüstü uygulamasıdır. Basit bir arayüz ile snippet'leri organize etmenize ve ihtiyaç duyduğunuzda hızlıca erişmenize olanak tanır.

---

## 🚀 Özellikler

- **Snippet Yönetimi**:
    - Yeni snippet ekleme, düzenleme ve silme işlemleri.
- **Tema Desteği**:
    - Varsayılan olarak **Karanlık Tema**.
    - Kullanıcı isteğine göre **Açık Tema** seçeneği.
- **Arama Fonksiyonu**:
    - Snippet başlıkları arasında hızlı arama yapabilme.
- **Kalıcı Depolama**:
    - Snippet'ler `snippets.json` dosyasında saklanır.
- **Kullanıcı Dostu Arayüz**:
    - Minimal ve anlaşılır bir tasarım ile hızlı erişim imkanı.

---

## 📂 Klasör Yapısı

```plaintext
FastCode/
├── main.go          # Ana uygulama dosyası
├── snippets.json    # Snippet'lerin saklandığı JSON dosyası
└── README.md        # Projenin açıklama dosyası
```

##  Kullanım Kılavuzu

- Go'nun bilgisayarınıza kurulu olduğundan emin olun:
  $ go version

- Proje klasörüne gidin:
  $ cd fast-code

- Uygulamayı başlatın:
  $ go run main.go

##  Özelliklerin Kullanımı

## Snippet Ekleme

1. "Başlık" alanına bir başlık yazın.
2. "İçerik" alanına kod snippet'inizi yazın.
3. "Kaydet" düğmesine tıklayarak snippet'i kaydedin.

###  Snippet Arama

1. Üstteki arama çubuğuna snippet başlığını yazın.
2. İlgili sonuçlar listede görünecektir.

### Snippet Düzenleme

1. Listeden düzenlemek istediğiniz bir snippet'e tıklayın.
2. Gerekli değişiklikleri yapıp "Kaydet" düğmesine tıklayın.

### Snippet Silme

1. Listeden silmek istediğiniz snippet'i seçin.
2. "Sil" düğmesine tıklayarak snippet'i kaldırın.

### Tema Değiştirme 

1. Tema seçim kutusundan "Karanlık" veya "Beyaz" temayı seçin.


## 🛠️ Kullanılan Teknolojiler

- Programlama Dili: Go (Golang)
- Arayüz Kütüphanesi: Fyne (https://fyne.io)


## 🤝 Katkı Sağlama

1. Bu repoyu forklayın:
   $ git clone https://github.com/Tahaknd/fast-code.git

2. Yeni bir branch oluşturun:
   $ git checkout -b yeni-ozellik

3. Değişiklik yapın ve commit edin:
   $ git commit -m "Yeni özellik eklendi: ..."

4. Değişikliklerinizi uzak repoya gönderin:
   $ git push origin yeni-ozellik

5. Pull Request gönderin.


## 🐞 Sorunlar ve Geri Bildirim 

- Bir sorunla karşılaşırsanız veya yeni bir özellik önermek isterseniz, bir issue oluşturabilirsiniz:
  https://github.com/Tahaknd/fast-code/issues

