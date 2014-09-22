#Un’introduzione a BDD

Traduzione in italiano dell’articolo Introducing BDD di Dan North


*Storia*: questo articolo è stato inizialmente pubblicato sulla rivista “Better Software” nel marzo del 2006. È stato tradotto in giapponese da Yukei Wachi e in coreano da HogJoo Lee.

Avevo un problema.
Nel periodo in cui utilizzavo e insegnavo tecniche agili come il Test-driven development (TDD) applicate a progetti nei più disparati ambienti, notavo che emergevano sempre gli stessi dubbi e le stesse incomprensioni. Gli sviluppatori volevano sapere da dove iniziare, cosa collaudare e cosa non collaudare, quanto collaudare in un colpo solo, come chiamare i propri test e come capire il perché un proprio test stesse fallendo.

Più mi addentravo in TDD più mi sembrava che il mio stesso viaggio fosse una serie di vicoli ciechi invece che un processo di graduale raffinamento, di “togli la cera, metti la cera”. Ricordo che pensavo più spesso “Se solo qualcuno me lo avesse detto prima” di quanto pensassi “Wow, si è aperta una porta”. Decisi che dovesse essere possibile presentare TDD in una forma in grado di condurre direttamente ai risultati buoni evitando tutti i trabocchetti.

La mia risposta è stata il behaviour-driven development (BDD). BDD deriva da comprovate pratiche agili ed è stato disegnato perché i team che si siano avvicinati da poco alle tecniche agili le trovino più accessibili ed efficaci. Nel tempo BDD è cresciuto fino ad abbracciare un quadro più ampio che include l’analisi agile e i test di accettazione automatici.

##I nomi dei test dovrebbero essere delle frasi

La prima occasione in cui ho esclamato “Aha!” è avvenuta quando mi è stato mostrato un tool, ingannevolmente semplice, chiamato agiledox e realizzato dal mio collega Chris Stevenson. Questo tool prende una classe JUnit e ne stampa i nomi dei metodi come se fossero ordinarie frasi, in modo che una classe di test simile a:

	public class CustomerLookupTest extends TestCase {
	    testFindsCustomerById() {
	        ...
	    }
	    testFailsForDuplicateCustomers() {
	        ...
	    }
	    ...
	}

restituisca qualcosa come:

	CustomerLookup
	- finds customer by id
	- fails for duplicate customers
	- ...

Il termine “test” viene omesso sia dal nome della classe che dal nome dei metodi e il nome del metodo viene convertito da camel-case a testo regolare. Questo è tutto quel che il tool fa, ma il risultato è sorprendente.

Gli sviluppatori scoprirono che il tool poteva essere utile per scrivere almeno un po’ della documentazione automatica, così iniziarono a scegliere per i nomi dei metodi delle vere frasi.
C’è di più: scoprirono che quando scrivevano i nomi dei metodi utilizzando il linguaggio del dominio del business, la documentazione generata risultava comprensibile agli esperti di business, agli analisti e ai tester.

##Un semplice modello per le frasi fa sì che i metodi siano ben a fuoco

È così che sono giunto alla convenzione di iniziare i nomi dei metodi dei miei testi con la parola should (dovrebbe). Questo modello per il nome del metodo – “La classe dovrebbe fare qualcosa” – aiuta a ricordare che si dovrebbero definire test solo per la classe presa in considerazione, in modo che risulti più facile concentrarsi. Quando capita di scoprire che si sta scrivendo un test il cui nome non soddisfi questo modello, si può sospettare che il comportamento possa appartenere ad altri contesti.

Per esempio, stavo scrivendo una classe per convalidare gli input provenienti dallo schermo. La maggior parte dei campi era costituito da ordinari dettagli del cliente (nome, cognome etc) ma c’era anche un campo per la data di nascita ed uno per l’età. Ero partito con lo scrivere una classe ClientDetailsValidatorTest con metodi come testShouldFailForMissingSurname e testShouldFailForMissingTitle.

Dopo di che, ho proceduto con il calcolo dell’età ed ho inserito una gran quantità di regole di business piuttosto rognose: come comportarsi se l’età e la data di nascita venissero fornite ma non dovessero coincidere? Cosa fare se la data di compleanno fosse il giorno corrente? Come avrei dovuto calcolare l’età se fosse fornita solo la data di nascita? Stavo progressivamente scrivendo ingombranti nomi di metodi di test per descrivere questo comportamento, così ho preso in considerazione l’idea spostare tutto da qualche altra parte. Questo mi portò a introdurre una nuova classe che chiamai AgeCalculator, con il suo AgeCalculatorTest. Avevo spostato tutto il comportamento relativo al calcolo dell’età nella nuova classe, così il validatore finiva per aver bisogno solo di un test relativo al calcolo dell’età per assicurarsi che l’interazione con il calcolatore venisse svolta correttamente.

Quando una classe svolge più di una singolo compito specifico, di solito la prendo come l’indicazione che dovrei introdurre altre classi e delegare loro un po’ del suo lavoro. Definisco il nuovo servizio in termini di interfaccia, descrivendo quel che fa e poi lo passo alla classe attraverso il suo costruttore:

	public class ClientDetailsValidator {
	 
	    private final AgeCalculator ageCalc;
	 
	    public ClientDetailsValidator(AgeCalculator ageCalc) {
	        this.ageCalc = ageCalc;
	    }
	}

Questo approccio nel collegare gli oggetti tra loro, noto come dependency injection, è particolarmente utile quando viene utilizzato insieme ai mock.

##Un nome espressivo per un test è utile quando il test fallisce

Col tempo mi accorsi che se modificando il codice provocavo il fallimento di un test, potevo leggere il nome del metodo e identificare il comportamento atteso del codice. Tipicamente accadeva una delle seguenti tre cose:

Avevo introdotto un bug. Colpa mia. Soluzione: correggere il bug
Il comportamento atteso risultava ancora pertinente ma era stato spostato altrove. Soluzione: sportare il test ed eventualmente modificarlo
Il comportamento non era più corretto, le premesse del sistema erano cambiato. Soluzione: eliminare il test
L’ultimo caso può capitare nei progetti agili, proprio perché la comprensione del problema evolve progressivamente nel tempo. Sfortunatamente i principianti del TDD hanno una paura innata a eliminare i test poiché ritengono che in qualche modo questo possa ridurre la qualità del proprio codice.

Un effetto più sottile della parola should (dovrebbe) diviene evidente non appena la si confronta con le più formali alternative will e shall (farà). Should implicitamente consente di mettere in discussione le premesse del test: “Dovrebbe? Dovrebbe davvero?” L’uso di should rende più semplice decidere se un test stia fallendo a causa di un bug appena introdotto o semplicemente perché le precedenti assunzioni a proposito del comportamento del sistema non siano più corrette.

**Behaviour** (comportamento) è una parola più utile di test. Adesso posseggo un tool, **agiledox**, per eliminare la parola test e un modello per i nomi di ogni metodo di test.

Mi parve subito chiaro che le difficoltà delle persone nel comprendere TDD derivassero dall’uso della parola test.
Non intendo dire che l’attività di test non sia intrinseca nel TDD: l’insieme risultante di metodi di test è un modo efficace per assicurare il funzionamento del proprio lavoro. Tuttavia, se i metodi non descrivono dettagliatamento il comportamento del sistema, allora ci stanno ingannevolmente cullando in una falsa sensazione di sicurezza.

Ho iniziato ad utilizzare la parola behaviour (comportamento) invece che test nelle mie questioni di TDD e non solo ho trovato che il termine calzasse ma anche che un’intera categorie di coaching question magicamente scompariva. Improvvisamente trovavo le risposte ad alcune di queste domande. Come scegliere il nome dei metodi di test diveniva semplice: si sceglie una frase che descriva il prossimo comportamento nel quale si è interessati. In che misura collaudare diveniva argomento questione di lana caprina: una singola frase deve descrivere un singolo comportamento. Quando un test falliva, semplicemente ripercorrevo il processo prima descritto: o avevo introdotto un bug o il comportamento era stato spostato o il test non era più significativo.

Passare dal pensare in termini di test a pensare in termini di comportamento è un moto di rivoluzione così profondo che ho iniziato a riferirmi al TDD come BDD, o behaviour-driven development.

##JBehave evidenzia il comportamento rispetto all’attività di test

Alla fine del 2003 ho deciso che fosse giunto il tempo di investire un po’ di tempo per concretizzare le mie idee. Ho iniziato a scrivere un tool sostitutivo per JUnit, JBehave, nel quale ho rimosso ogni riferimento al concetto di test rimpiazzandolo con un vocabolario costruito intorno al concetto di verifica di comportamento. Ho voluto provare per vedere quanto un framework potesse evolvere aderendo fedelmente al mio nuovo mantra del behaviour-driven. Ho anche pensato che sarebbe stato un valido strumento di insegnamento per introdurre TDD e BDD senza la distrazione di un vocabolario costruito intorno ai termini del mondo dei test.

Per definire il comportamento di un un’ipotetica classe CustomerLookup avrei scritto una classe chiamata, per esempio, CustomerLookupBehaviour. Avrebbe contenuto metodi il cui nome sarebbe iniziato con la parola should. Il tool di esecuzione del behaviour avrebbe istanziato la classe di behaviour ed avrebbe invocato in sequenza ogni metodo di behaviour, proprio come JUnit faceva per i metodi di test. Il tool avrebbe dovuto prodotto un report sullo stato di avanzamento ed un sommario finale.

Il mio primo traguardo era quello di rendere JBehave capace di auto-verificarsi. Non avevo che da aggiungere un behavior capace di lanciare se stesso. Alla fine, sono stato in grado di migrare tutti i test JUnit in behaviour di JBehave e di ottenere feedback immediati proprio come in JUnit.

##Individuare il successivo behavior in ordine di importanza

Poi ho scoperto il concetto di valore di business. Naturalmente ero sempre stato consapevole che stessi scrivendo software per una ragione, ma non avevo mai pensato seriamente al valore del codice che realizzato prima di allora. Un mio collega, l’analista di business Chriss Matts, mi ha consentito di pensare al valore di business nel contesto del behaviour-driven development.

Con la premessa che il mio obiettivo principale fosse quello di rendere JBehave auto-contenuto, scoprii che quel che mi aiutava a stare concentrato era il chiedermi: qual è la prossima cosa importante che il sistema non fa?

Questa domanda richiede di individuare il valore delle feature ancora non implementate e di assegnare loro una priorità. Facilita anche la formulazione dei nomi dei metodi di behaviour: il sistema non fa X (dove X è un qualche comportamento significativo) e X è importante, il che significa che il sistema dovrebbe fare X; così, il prossimo behaviour è semplicemente:

	public void shouldDoX() {
	    // ...
	}

Ora avevo la risposta a un nuovo quesito sul TDD, cioè da dove partire.

##Anche i requisiti sono behaviour

A questo punto, disponevo di un framework che mi aiutava a capire – e, più importante ancora, a spiegare – come funzionasse il TDD e di un approccio che mi permettesse di evitare tutti i tranelli che avevo sempre incontrato.

Verso la fine del 2004, mentre stavo descrivendo il mio nuovo vocabolario basato sul comportamento a Matts, lui mi disse “È esattamente come una fase di analisi”. Ci fu una lunga pausa mentre ragionammo su quanto appena detto, dopo di che decidemmo di applicare l’intero pensiero behaviour-driven alla definizione dei requisiti.

Se fossimo stati in grado di sviluppare un vocabolario sistematico per gli analisti, i tester, gli sviluppatori e i responsabili del business saremmo stati sulla buona strada per eliminare alcune delle ambiguità e delle incomprensioni che emergono quando i tecnici cercano di parlare con i responsabili del business.

##Il BDD fornisce un “ubiquitous language” per l’analisi

Eric Evans ha pubblicato da poco il suo best-seller Domain-Driven Design, nel quale descrive il concetto di modellazione di un sistema per mezzo di un “ubiquitous language”, un linguaggio legato così strettamente sul dominio di business che il codice del software viene permeato dallo stesso vocabolario del mondo business.

Io e Chris abbiamo compreso che stavano tentando esattamente di definire un ubiquitous language per il processo di analisi! Stavamo partendo col piede giusto.
Nella nostra compagnia veniva convenzionalmente utilizzato un modello per le story simile a questo:

	Come [X]
	Voglio [Y]
	Così che [Z]

dove Y è una qualche feature, Z il beneficio ottenuto o il valore della feature e X la persona (o il ruolo) che avrebbe goduto del vantaggio.

La forza di questo modello risiede nel fatto che costringe ad identificare il valore che possiede il consegnare in produzione una story nel momento stesso in cui la si definisce la prima volta. Quando non esiste valore di business dietro una story, l’applicazione del modello fa ottenere qualcosa di simile a “Io voglio [una qualche funzionalità] affinche [la voglio e basta, ok?]“. Questo può facilitare il ridimensionamento di alcuni dei requisiti più esoterici.

Dal punto in cui stavamo partendo, io e Matt ci stavamo accingendo a scoprire quel che ogni tester del mondo agile già sapeva: il behaviour di una story è semplicemente il suo criterio di accettazione – se il sistema soddisfa tutti i i criteri di accettazione si comporta correttamente; se non lo fa, non si comporta correttamente. Così, creammo un modello per catturare il criterio di accettazione di una storia.

Il modello avrebbe dovuto essere sufficientemente libero da non apparire artificiale e da non limitare l’azione degli analisti, ma al contempo sufficientemente strutturato da permetterci di decomporre la storia nei suoi elementi costituenti e di automatizzarli. Partimmo col descrivere il criterio di accettazione in termini di scenario, il che ci portò alla seguente forma:

	Dato un qualche contesto iniziale, ("given")
	Quando accade qualcosa, ("when")
	allora assicurati che succeda qualcos'altro. ("then")

Per andare sul concreto, usiamo il classico esempio di un bancomat. Una tipica story card potrebbe assomigliare a

+Titolo: un cliente preleva del denaro+
Come cliente
voglio poter prelevare del contante dal bancomat
in modo da poter evitare la fila alla banca.
Ora, come fare a sapere quando la storia sia completamente implementata? Ci sono vari scenari da prendere in considerazione: il conto potrebbe avere del credito oppure potrebbe essere scoperto ma entro i limiti di scoperto, oppure potrebbe essere scoperto e fuori dai limiti. Naturalmente, ci sono anche altri scenari, come il caso di conto in attivo ma di richiesta di prelievo così alta da portarlo allo scoperto, oppure quello di erogatore privo di sufficiente contante.

Ricorrendo al modello “Dato…quando…allora…” (given-when-then) i primi due scenari apparirebbero così:

+Scenario 1: Il conto è in attivo
Dato un conto corrente in attivo
E una carta valida
E dato un bancomat con sufficiente contante
Quando il cliente esegue un prelievo
Allora il contante viene addebitato sul conto
E il denaro viene erogato
Si noti l’uso delle congiunzioni e per collegare tra loro più clausole dato (given) o più clausole allora (then), come si farebbe nel linguaggio naturale.

	+Scenario 2: Il conto è scoperto oltre i limiti consentiti+
	Dato un conto scoprto
	Ed una carta valida
	Quando il cliente richiede un prelievlo
	Allora un messaggio di rifiuto viene visualizzato
	E il denaro non viene erogato
	E la carta viene restituita.
Entrambi gli scenari sono basati sui medesimi eventi e condividono perfino alcuni given e when. Vogliamo trarre vantaggio da questo in modo da poter fare riuso di “*given*”, eventi e “*then*”.

##I test di accettazione dovrebbero essere eseguibili

I frammenti dello scenario — cioè i given, i when e i then — sono ad un livello di dettaglio abbastanza fine da poter essere espressi attraverso il codice. JBehave definisce un modello ad oggetti che permette di mappare direttamente lo scenario in classi Java.

Si scrive una classe per rappresentare ogni given:

	public class AccountIsInCredit implements Given {
	    public void setup(World world) {
	        ...
	    }
	}
	public class CardIsValid implements Given {
	    public void setup(World world) {
	        ...
	    }
	}

ed una per la condizione when

	public class CustomerRequestsCash implements Event {
	    public void occurIn(World world) {
	        ...
	    }
	}

e così via per le clausole then. JBehave colega tutto quanto insieme ed esegue il risultato. Crea un “mondo”, ovvero uno spazio nel quale conservare i propri oggetti, e lo fornisce a disposizione di ogni given in modo che questi possano popolarlo con degli stati noti. Dopo di che JBehave chiede all’evento when di “accadere” nel “mondo”, il che porta al vero e proprio comportamento dello scenario. Infine, il controllo viene passato ai vari then definiti nella story.

Il fatto di utilizzare una classe separata per rappresentare ogni frammento permette di riutilizzare i frammenti in altri scenari o in altre story. All’inizio, i frammenti vengono implementati con dei mock che assicurino che un conto sia in attivo o che una carta si avalida. Saranno i punti di partenza per l’implementazione del behaviour. Via via che l’implementazione dell’applicazione procede il codice dietro ai given e ai then viene modificato in modo che utilizzi le classi che sono state implementate, così che quando lo scenario sia stato completato si siano trasformati in veri e propri test funzionali end-to-end.

##Il presente e il futuro di BDD

Dopo un breve periodo di inattività, è ripreso pienamente lo sviluppo di JBehave. Il core è abbastanza completo e robusto. Il prossimo passo è l0integrazione con gli IDE più popolari per Java come IntelliJ IDEA ed Eclipse.

Dave Astels sta attivamente promuovendo BDD. Il suo blog e i numerosi articoli che ha pubblicato hanno prodotto una valanga di attività, in special modo il progetto rspec pensato per fornire un framework BDD per Ruby. Io ho iniziato a lavorare a rbehave che sarà un’implementazione di JBEhave per Ruby.

Alcuni dei miei colleghi hanno iniziato ad utilizzare le tecniche BDD su una gran varietà di progetti del mondo reale e le stanno trovando efficaci. Lo story runner di JBehave, la parte che verifica i criteri di accettazione, è tutt’ora in sviluppo.

La prospettiva è quella di ottenere un editor “andata e ritorno” in modo che gli analisti di Business ed i tester possano catturare le story mediante un ordinario editor di testi in grado di generare gli stub per le classi behaviour, utilizzando solo il linguaggio di del dominio del business. BDD sta evolvendo con l’aiuto di molte persone ed io sono profondamente grado a ognuna di loro.